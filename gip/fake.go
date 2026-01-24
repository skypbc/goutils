package gip

import (
	"crypto/rand"
	"encoding/binary"
	"net"
	"net/netip"
)

// randomUint32 возвращает случайное uint32
func randomUint32() uint32 {
	var b [4]byte
	_, _ = rand.Read(b[:])
	return binary.BigEndian.Uint32(b[:])
}

// randomUint128 возвращает случайное 128-битное число как [16]byte
func randomUint128() [16]byte {
	var b [16]byte
	_, _ = rand.Read(b[:])
	return b
}

// FakeIPv4 генерирует случайный публичный IPv4
func FakeIPv4() string {
	for {
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, randomUint32())

		addr, ok := netip.AddrFromSlice(ip)
		if !ok {
			continue
		}

		if addr.IsPrivate() || addr.IsLoopback() || addr.IsMulticast() ||
			addr.IsLinkLocalUnicast() || addr.IsLinkLocalMulticast() {
			continue
		}

		return addr.String()
	}
}

// FakeIPv6 генерирует случайный публичный IPv6
func FakeIPv6() string {
	for {
		raw := randomUint128()
		ip := net.IP(raw[:])

		addr, ok := netip.AddrFromSlice(ip)
		if !ok {
			continue
		}

		if addr.IsPrivate() || addr.IsLoopback() || addr.IsMulticast() ||
			addr.IsLinkLocalUnicast() || addr.IsLinkLocalMulticast() {
			continue
		}

		// В Go нет IsSiteLocal, но эти адреса — устаревшие (FEC0::/10).
		if addr.Is6() && addr.As16()[0]&0xfe == 0xfc {
			// fc00::/7 (ULA, "unique local")
			continue
		}
		if addr.Is6() && (addr.As16()[0]&0xfe == 0xfe) && (addr.As16()[1]&0xc0 == 0xc0) {
			// fec0::/10 (site-local, deprecated)
			continue
		}

		return addr.String()
	}
}
