package ssl

import (
	. "github.com/pascalgaut/filestash/server/common"
)

var (
	keyPEMPath  func() string
	certPEMPath func() string
)

func init() {
	keyPEMPath = func() string {
		return GetAbsolutePath(CERT_PATH, "key.pem")
	}
	certPEMPath = func() string {
		return GetAbsolutePath(CERT_PATH, "cert.pem")
	}
}

func Clear() {
	clearPrivateKey()
	clearCert()
}
