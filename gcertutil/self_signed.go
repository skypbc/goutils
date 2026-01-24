package gcertutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"net"
	"sort"
	"strings"
	"time"
)

// GenerateRootCARSA создаёт самоподписанный корневой CA сертификат с RSA-2048
func GenerateRootCARSA(commonName string, validFor time.Duration) (caCertPEM, caKeyPEM []byte, err error) {
	if commonName = strings.TrimSpace(commonName); commonName == "" {
		return nil, nil, errors.New("commonName is required")
	}
	if validFor <= 0 {
		validFor = 10 * 365 * 24 * time.Hour // по умолчанию 10 лет
	}

	// генерируем приватный ключ RSA-2048
	caPriv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// серийный номер
	serial, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, nil, err
	}

	// вычисляем SubjectKeyId для CA
	pubDER, err := x509.MarshalPKIXPublicKey(&caPriv.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	ski := sha1.Sum(pubDER)

	notBefore := time.Now().Add(-1 * time.Hour)
	tpl := x509.Certificate{
		SerialNumber: serial,
		Subject: pkix.Name{
			CommonName:   commonName,
			Organization: []string{"Dev Local Root CA"},
		},
		NotBefore:             notBefore,
		NotAfter:              notBefore.Add(validFor),
		BasicConstraintsValid: true,

		// CA-флаги
		IsCA:               true,
		MaxPathLen:         0,
		MaxPathLenZero:     true,
		KeyUsage:           x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		SignatureAlgorithm: x509.SHA256WithRSA,
		SubjectKeyId:       ski[:],
		AuthorityKeyId:     ski[:], // у корня AKI = SKI
	}

	// самоподпись корневого сертификата
	der, err := x509.CreateCertificate(rand.Reader, &tpl, &tpl, &caPriv.PublicKey, caPriv)
	if err != nil {
		return nil, nil, err
	}

	// кодируем в PEM
	caCertPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
	caKeyPEM = pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(caPriv)})
	return caCertPEM, caKeyPEM, nil
}

// IssueServerCertRSA принимает PEM корневого CA и выпускает серверный сертификат.
// dnsNames — список доменных имён (включая wildcard), ipSans — список IP для SAN.
func IssueServerCertRSA(caCertPEM, caKeyPEM []byte, dnsNames []string, ipSans []net.IP, validFor time.Duration) (srvCertPEM, srvKeyPEM []byte, err error) {
	// парсим CA сертификат
	caCert, err := parseCertificateFromPEM(caCertPEM)
	if err != nil {
		return nil, nil, err
	}
	// парсим CA приватный ключ (поддержим PKCS#1 и PKCS#8)
	caKey, err := parseRSAPrivateKeyFromPEM(caKeyPEM)
	if err != nil {
		return nil, nil, err
	}

	// нормализуем SAN
	dns := uniqueNonEmpty(dnsNames)
	ips := uniqueIPs(ipSans)
	if len(dns) == 0 && len(ips) == 0 {
		return nil, nil, errors.New("at least one DNS name or IP must be provided")
	}
	if validFor <= 0 {
		validFor = 365 * 24 * time.Hour
	}

	// генерируем приватный ключ сервера
	srvPriv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// серийный номер для сервера
	serial, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, nil, err
	}

	// SKI для сервера
	pubDER, err := x509.MarshalPKIXPublicKey(&srvPriv.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	ski := sha1.Sum(pubDER)

	// AKI берём из SubjectKeyId CA (если есть)
	var aki []byte
	if len(caCert.SubjectKeyId) > 0 {
		aki = caCert.SubjectKeyId
	}

	// CN — первый домен, если он есть
	subject := pkix.Name{
		Organization: []string{"Dev Server Certificate"},
	}
	if len(dns) > 0 {
		subject.CommonName = dns[0]
	}

	notBefore := time.Now().Add(-1 * time.Hour)
	tpl := x509.Certificate{
		SerialNumber:          serial,
		Subject:               subject,
		NotBefore:             notBefore,
		NotAfter:              notBefore.Add(validFor),
		BasicConstraintsValid: true,

		// серверные флаги
		IsCA:               false,
		SignatureAlgorithm: x509.SHA256WithRSA,
		KeyUsage:           x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:        []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},

		// SAN
		DNSNames:    dns,
		IPAddresses: ips,

		SubjectKeyId:   ski[:],
		AuthorityKeyId: aki,
	}

	// подписываем серверный сертификат корневым CA
	der, err := x509.CreateCertificate(rand.Reader, &tpl, caCert, &srvPriv.PublicKey, caKey)
	if err != nil {
		return nil, nil, err
	}

	// кодируем в PEM
	srvCertPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
	srvKeyPEM = pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(srvPriv)})
	return srvCertPEM, srvKeyPEM, nil
}

func parseCertificateFromPEM(pemBytes []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, errors.New("failed to parse CERTIFICATE PEM")
	}
	return x509.ParseCertificate(block.Bytes)
}

func parseRSAPrivateKeyFromPEM(pemBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("failed to parse PRIVATE KEY PEM")
	}
	switch block.Type {
	case "RSA PRIVATE KEY": // PKCS#1
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	case "PRIVATE KEY": // PKCS#8
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rsaKey, ok := key.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("not an RSA private key")
		}
		return rsaKey, nil
	default:
		return nil, errors.New("unsupported private key type: " + block.Type)
	}
}

func uniqueNonEmpty(vals []string) []string {
	seen := make(map[string]struct{}, len(vals))
	out := make([]string, 0, len(vals))
	for _, v := range vals {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}
	// детерминированный порядок
	sort.Strings(out)
	return out
}

func uniqueIPs(ips []net.IP) []net.IP {
	seen := make(map[string]struct{}, len(ips))
	out := make([]net.IP, 0, len(ips))
	for _, ip := range ips {
		if ip == nil {
			continue
		}
		k := ip.String()
		if _, ok := seen[k]; !ok {
			seen[k] = struct{}{}
			out = append(out, ip)
		}
	}
	// детерминированный порядок
	sort.Slice(out, func(i, j int) bool { return out[i].String() < out[j].String() })
	return out
}
