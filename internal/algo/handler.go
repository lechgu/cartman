package algo

import (
	"crypto/x509"
)

type Handler interface {
	GenerateKeyPair() (KeyPair, error)
	SignatureAlgorithm() x509.SignatureAlgorithm
}
