package client

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func getScheme(ep string) string {
	// TODO: use url.Parse somehow?
	for _, scheme := range []string{"http", "https", "tcp", "grpc", "grpcs", "grpc+tls"} {
		if strings.HasPrefix(ep, scheme+"://") {
			return scheme
		}
	}
	// false-positive will lead to worse UX on misconfiguration, but it's more important
	// not to break any valid configuration.
	return ""
}

// If you are patching Perforator, you can initialize this in a `func init()`
var (
	defaultHost string = ""
)

const (
	defaultSecurePort   uint16 = 7618
	defaultInsecurePort uint16 = 7617
)

func parseEndpoint(ep string, secure *bool) (endpoint, error) {
	var e endpoint
	scheme := getScheme(ep)
	if scheme != "" {
		return e, fmt.Errorf("invalid endpoint: scheme must be unset, but got %q", scheme)
	}
	u, err := url.Parse("grpc://" + ep)
	if err != nil {
		return e, fmt.Errorf("failed to parse endpoint: %w", err)
	}
	e.host = u.Hostname()

	if secure != nil {
		e.secure = *secure
	} else {
		e.secure = true
	}

	if u.Port() != "" {
		var port uint64
		port, err = strconv.ParseUint(u.Port(), 10, 16)
		if err != nil {
			return e, fmt.Errorf("failed to parse port %q in endpoint: %w", u.Port(), err)
		}
		e.port = uint16(port)
	} else {
		if e.secure {
			e.port = defaultSecurePort
		} else {
			e.port = defaultInsecurePort
		}
	}

	return e, nil
}
