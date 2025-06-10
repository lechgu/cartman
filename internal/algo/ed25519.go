package algo

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
)

type ed25519Handler struct{}

func newEd25519() Handler {
	return &ed25519Handler{}
}

func (e *ed25519Handler) GenerateKeyPair() (KeyPair, error) {
	var keyPair KeyPair
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return keyPair, err
	}
	keyPair.PrivateKey = priv
	keyPair.PublicKey = pub
	return keyPair, nil
}

func (e *ed25519Handler) SignatureAlgorithm() x509.SignatureAlgorithm {
	return x509.PureEd25519
}
