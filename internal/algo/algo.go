package algo

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"fmt"
)

const (
	RSA2048  = "rsa2048"
	RSA3072  = "rsa3072"
	RSA4096  = "rsa4096"
	ECDSA256 = "ecdsa256"
	ECDSA384 = "ecdsa384"
	ECDSA521 = "ecdsa521"
	ED25519  = "ed25519"
)

var Supported = []string{
	RSA2048,
	RSA3072,
	RSA4096,
	ECDSA256,
	ECDSA384,
	ECDSA521,
	ED25519,
}

func NewHandler(which string) (Handler, error) {
	switch which {
	case RSA2048:
		return newRSA(2048)
	case RSA3072:
		return newRSA(3072)
	case RSA4096:
		return newRSA(4096)
	case ECDSA256:
		return newECDSA(256)
	case ECDSA384:
		return newECDSA(384)
	case ECDSA521:
		return newECDSA(521)
	case ED25519:
		return newEd25519(), nil
	default:
		return nil, nil
	}
}

func MatchHandler(priv crypto.PrivateKey) (Handler, error) {
	switch key := priv.(type) {
	case *rsa.PrivateKey:
		bits := key.N.BitLen()
		switch bits {
		case 2048, 3072, 4096:
			return newRSA(bits)
		default:
			return nil, fmt.Errorf("unsupported RSA key size: %d", bits)
		}

	case *ecdsa.PrivateKey:
		bits := key.Params().BitSize
		return newECDSA(bits)

	case ed25519.PrivateKey:
		return newEd25519(), nil

	default:
		return nil, fmt.Errorf("unsupported private key type: %T", priv)
	}
}
