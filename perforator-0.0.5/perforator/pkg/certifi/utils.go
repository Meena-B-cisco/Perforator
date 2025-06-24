package certifi

import (
	"crypto/x509"
	"fmt"
	"os"
)

// CertPoolFromFile creates certificate pool from the specified file path.
// If no path is provided, it returns the default certificate pool.
func CertPoolFromFile(caFile string) (*x509.CertPool, error) {
	if caFile != "" {
		caCert, err := os.ReadFile(caFile)
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM(caCert) {
			return nil, fmt.Errorf("failed to add certificate: %s", caFile)
		}
		return caCertPool, nil
	}

	return NewDefaultCertPool()
}
