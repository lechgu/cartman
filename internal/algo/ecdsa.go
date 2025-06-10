package algo

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"fmt"
)

type ecdsaHandler struct {
	curve elliptic.Curve
	algo  x509.SignatureAlgorithm
}

func newECDSA(bits int) (Handler, error) {
	switch bits {
	case 256:
		return &ecdsaHandler{
			curve: elliptic.P256(),
			algo:  x509.ECDSAWithSHA256,
		}, nil
	case 384:
		return &ecdsaHandler{
			curve: elliptic.P384(),
			algo:  x509.ECDSAWithSHA384,
		}, nil
	case 521:
		return &ecdsaHandler{
			curve: elliptic.P521(),
			algo:  x509.ECDSAWithSHA512,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported ECDSA size: %d (only 256, 384, or 521 allowed)", bits)
	}
}

func (e *ecdsaHandler) GenerateKeyPair() (KeyPair, error) {
	var keyPair KeyPair
	priv, err := ecdsa.GenerateKey(e.curve, rand.Reader)
	if err != nil {
		return keyPair, err
	}
	keyPair.PrivateKey = priv
	keyPair.PublicKey = priv.Public()
	return keyPair, nil
}

func (e *ecdsaHandler) SignatureAlgorithm() x509.SignatureAlgorithm {
	return e.algo
}
