package certificates

import (
	"crypto"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"net"
	"time"

	"github.com/lechgu/cartman/internal/algo"
)

func InitRoot(h algo.Handler, keyPair algo.KeyPair, validityDays int, subject *pkix.Name) (*x509.Certificate, error) {
	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, err
	}

	nbf := time.Now().Add(-5 * time.Minute)
	exp := nbf.AddDate(0, 0, validityDays)

	template := &x509.Certificate{
		SerialNumber:          serialNumber,
		Subject:               *subject,
		NotBefore:             nbf,
		NotAfter:              exp,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            0,
		MaxPathLenZero:        true,
		SignatureAlgorithm:    h.SignatureAlgorithm(),
	}

	certDER, err := x509.CreateCertificate(
		rand.Reader,
		template,
		template,
		keyPair.PublicKey,
		keyPair.PrivateKey,
	)
	if err != nil {
		return nil, err
	}

	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

func Issue(h algo.Handler, signer *algo.Signer, validityDays int, subject *pkix.Name, dnsNames []string, ipAddrs []net.IP) (*x509.Certificate, crypto.PrivateKey, error) {
	keyPair, err := h.GenerateKeyPair()
	if err != nil {
		return nil, nil, err
	}
	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, nil, err
	}

	nbf := time.Now().Add(-5 * time.Minute)
	exp := nbf.AddDate(0, 0, validityDays)

	template := &x509.Certificate{
		SerialNumber:          serialNumber,
		Subject:               *subject,
		NotBefore:             nbf,
		NotAfter:              exp,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
		DNSNames:              dnsNames,
		IPAddresses:           ipAddrs,
		SignatureAlgorithm:    h.SignatureAlgorithm(),
	}

	certDER, err := x509.CreateCertificate(
		rand.Reader,
		template,
		signer.Cert,
		keyPair.PublicKey,
		signer.Key,
	)
	if err != nil {
		return nil, nil, err
	}

	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		return nil, nil, err
	}
	return cert, keyPair.PrivateKey, nil
}
