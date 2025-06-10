package algo

import (
	"crypto"
	"crypto/x509"
)

type Signer struct {
	Cert *x509.Certificate
	Key  crypto.PrivateKey
}
