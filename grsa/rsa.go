package grsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/skypbc/goutils/gerrors"
	"os"
	"strings"
)

type Key struct {
	headers map[string]string
	type_   string
	priv    *rsa.PrivateKey
	pub     *rsa.PublicKey
}

func (k *Key) Headers() map[string]string {
	return k.headers
}

func (k *Key) Type() string {
	return k.type_
}

func (k *Key) Priv() *rsa.PrivateKey {
	return k.priv
}

func (k *Key) Pub() *rsa.PublicKey {
	if k.pub == nil && k.priv != nil {
		return &k.priv.PublicKey
	}
	return k.pub
}

func PubKeyToBytes(key *rsa.PublicKey, headers ...map[string]string) (bytes []byte, err error) {
	if key == nil {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Private key is nil")
	}
	pubBytes := x509.MarshalPKCS1PublicKey(key)
	pubBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubBytes,
	}
	if len(headers) > 0 {
		pubBlock.Headers = headers[0]
	}
	return pem.EncodeToMemory(pubBlock), nil
}

func PubKeyToRawBytes(key *rsa.PublicKey) (bytes []byte, err error) {
	if key == nil {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Public key is nil")
	}
	pubBytes := x509.MarshalPKCS1PublicKey(key)
	if len(pubBytes) == 0 {
		return nil, gerrors.NewUnknownError().
			SetTemplate("Failed to convert RSA public key to raw bytes")
	}
	return pubBytes, nil
}

func PubKeyToString(key *rsa.PublicKey, headers ...map[string]string) (str string, err error) {
	pubBytes, err := PubKeyToBytes(key, headers...)
	if err != nil {
		return "", err
	}
	return string(pubBytes), nil
}

func PubKeyToRawString(key *rsa.PublicKey) (str string, err error) {
	if key == nil {
		return "", gerrors.NewIncorrectParamsError().
			SetTemplate("Public key is nil")
	}
	pubBytes := x509.MarshalPKCS1PublicKey(key)
	str = base64.StdEncoding.EncodeToString(pubBytes)
	if len(str) == 0 {
		return "", gerrors.NewUnknownError().
			SetTemplate("Failed to convert RSA public key to raw string")
	}
	return str, nil
}

func PubKeyToFile(path string, key *rsa.PublicKey, headers ...map[string]string) (err error) {
	if key == nil {
		return gerrors.NewIncorrectParamsError().
			SetTemplate("Public key is nil")
	}
	pubBytes, err := PubKeyToBytes(key, headers...)
	if err != nil {
		return err
	}
	if err = os.WriteFile(path, pubBytes, 0644); err != nil {
		return gerrors.NewFileWriteError(err).
			SetTemplate(`Failed to write public key to "{path}" file`).
			AddStr("path", path)
	}
	return nil
}

func PubKeyFromBytes(bytes []byte) (key *Key, err error) {
	if len(bytes) == 0 {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Public key bytes are empty")
	}
	block, _ := pem.Decode(bytes)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, gerrors.NewParseError().
			SetTemplate("Invalid PEM block for RSA public key")
	}
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, gerrors.NewParseError(err).
			SetTemplate("Failed to parse RSA public key")
	}
	return &Key{
		headers: block.Headers,
		type_:   block.Type,
		pub:     pubKey,
	}, nil
}

func PubKeyFromRawBytes(bytes []byte) (key *Key, err error) {
	if len(bytes) == 0 {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Public key bytes are empty")
	}
	pubKey, err := x509.ParsePKCS1PublicKey(bytes)
	if err != nil {
		return nil, gerrors.NewParseError(err).
			SetTemplate("Failed to parse RSA public key")
	}
	return &Key{
		pub: pubKey,
	}, nil
}

func PubKeyFromString(str string) (key *Key, err error) {
	if len(str) == 0 {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Public key string is empty")
	}
	return PubKeyFromBytes([]byte(str))
}

func PubKeyFromFile(path string) (key *Key, err error) {
	if len(path) == 0 {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Path to public key file is empty")
	}
	pubBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, gerrors.NewFileReadError(err).
			SetTemplate(`Failed to read public key from "{path}" file`).
			AddStr("path", path)
	}
	return PubKeyFromBytes(pubBytes)
}

func PrivKeyToBytes(key *rsa.PrivateKey, headers ...map[string]string) (bytes []byte, err error) {
	if key == nil {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Private key is nil")
	}
	privBytes := x509.MarshalPKCS1PrivateKey(key)
	privBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	}
	if len(headers) > 0 {
		privBlock.Headers = headers[0]
	}
	return pem.EncodeToMemory(privBlock), nil
}

func PrivKeyToRawBytes(key *rsa.PrivateKey) (bytes []byte, err error) {
	if key == nil {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Private key is nil")
	}
	privBytes := x509.MarshalPKCS1PrivateKey(key)
	if len(privBytes) == 0 {
		return nil, gerrors.NewUnknownError().
			SetTemplate("Failed to convert RSA private key to raw bytes")
	}
	return privBytes, nil
}

func PrivKeyToString(key *rsa.PrivateKey, headers ...map[string]string) (str string, err error) {
	privBytes, err := PrivKeyToBytes(key, headers...)
	if err != nil {
		return "", err
	}
	return string(privBytes), nil
}

func PrivKeyToRawString(key *rsa.PrivateKey) (str string, err error) {
	if key == nil {
		return "", gerrors.NewIncorrectParamsError().
			SetTemplate("Private key is nil")
	}
	privBytes := x509.MarshalPKCS1PrivateKey(key)
	str = base64.StdEncoding.EncodeToString(privBytes)
	if len(str) == 0 {
		return "", gerrors.NewUnknownError().
			SetTemplate("Failed to convert RSA private key to raw string")
	}
	return str, nil
}

func PrivKeyToFile(path string, key *rsa.PrivateKey, headers ...map[string]string) (err error) {
	if key == nil {
		return gerrors.NewIncorrectParamsError().
			SetTemplate("Private key is nil")
	}
	privBytes, err := PrivKeyToBytes(key, headers...)
	if err != nil {
		return err
	}
	if err = os.WriteFile(path, privBytes, 0600); err != nil {
		return gerrors.NewFileWriteError(err).
			SetTemplate(`Failed to write private key to "{path}" file`).
			AddStr("path", path)
	}
	return nil
}

func PrivKeyFromBytes(bytes []byte) (key *Key, err error) {
	if len(bytes) == 0 {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Private key bytes are empty")
	}
	block, _ := pem.Decode(bytes)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, gerrors.NewParseError().
			SetTemplate("Invalid PEM block for RSA private key")
	}
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, gerrors.NewParseError(err).
			SetTemplate("Failed to parse RSA private key")
	}
	return &Key{
		headers: block.Headers,
		type_:   block.Type,
		priv:    privKey,
	}, nil
}

func PrivKeyFromRawBytes(bytes []byte) (key *Key, err error) {
	if len(bytes) == 0 {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Private key bytes are empty")
	}
	privKey, err := x509.ParsePKCS1PrivateKey(bytes)
	if err != nil {
		return nil, gerrors.NewParseError(err).
			SetTemplate("Failed to parse RSA private key")
	}
	return &Key{
		priv: privKey,
	}, nil
}

func PrivKeyFromString(str string) (key *Key, err error) {
	if len(str) == 0 {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Private key string is empty")
	}
	return PrivKeyFromBytes([]byte(str))
}

func PrivKeyFromRawString(str string) (key *Key, err error) {
	if len(str) == 0 {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Private key string is empty")
	}
	str = strings.Join(strings.Split(strings.TrimSpace(str), "\n"), "")
	if len(str) == 0 {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Private key string is empty after trimming")
	}
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, gerrors.NewParseError(err).
			SetTemplate("Failed to decode key")
	}
	return PrivKeyFromRawBytes(data)
}

func PrivKeyFromFile(path string) (key *Key, err error) {
	if len(path) == 0 {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Path to private key file is empty")
	}
	privBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, gerrors.NewFileReadError(err).
			SetTemplate(`Failed to read private key from "{path}" file`).
			AddStr("path", path)
	}
	return PrivKeyFromBytes(privBytes)
}

type CreateOpts struct {
	Size    int
	Headers map[string]string
}

func PrivKeyCreate(opts ...CreateOpts) (key *Key, err error) {
	keySize := 2048
	if len(opts) > 0 {
		if opts[0].Size > 0 {
			keySize = opts[0].Size
		}
	}
	if keySize < 1024 {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Key size must be at least 1024 bits, got {key_size}").
			AddInt("key_size", keySize)
	}
	if keySize > 4096 {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate("Key size must not exceed 4096 bits, got {key_size}").
			AddInt("key_size", keySize)
	}
	privKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return nil, gerrors.Wrap(err).
			SetTemplate("Failed to generate RSA private key")
	}
	key = &Key{
		priv: privKey,
	}
	if len(opts) > 0 && opts[0].Headers != nil {
		key.headers = opts[0].Headers
	}
	return key, nil
}
