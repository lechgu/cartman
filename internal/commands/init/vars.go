package init

import "crypto/x509/pkix"

var (
	algorithm    string
	force        bool
	validityDays int
	subject      pkix.Name
)
