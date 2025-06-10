package issue

import (
	"crypto/x509/pkix"
	"net"
)

var (
	force        bool
	validityDays int
	subject      pkix.Name
	dnsNames     []string
	ipAddresses  []net.IP
)
