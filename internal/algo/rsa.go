package algo

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
)

type rsaHandler struct {
	bits int
}

func newRSA(bits int) (Handler, error) {
	switch bits {
	case 2048, 3072, 4096:
		return &rsaHandler{bits: bits}, nil
	default:
		return nil, fmt.Errorf("unsupported RSA key size: %d (only 2048, 3072, or 4096 allowed)", bits)
	}
}

func (r *rsaHandler) GenerateKeyPair() (KeyPair, error) {
	var keyPair KeyPair
	priv, err := rsa.GenerateKey(rand.Reader, r.bits)
	if err != nil {
		return keyPair, err
	}
	keyPair.PrivateKey = priv
	keyPair.PublicKey = priv.Public()
	return keyPair, nil
}

func (r *rsaHandler) SignatureAlgorithm() x509.SignatureAlgorithm {
	switch {
	case r.bits >= 4096:
		return x509.SHA512WithRSA
	case r.bits >= 3072:
		return x509.SHA384WithRSA
	default:
		return x509.SHA256WithRSA
	}
}
