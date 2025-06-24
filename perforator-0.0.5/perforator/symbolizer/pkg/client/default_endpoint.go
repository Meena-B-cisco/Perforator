package client

import (
	"fmt"
	"os"
	"strconv"
)

const perforatorEndpointEnv = `PERFORATOR_ENDPOINT`
const perforatorSecurityLevelEnv = `PERFORATOR_SECURE`

func getDefaultPerforatorEndpoint() (endpoint, error) {
	var ep string
	ep, ok := os.LookupEnv(perforatorEndpointEnv)
	if !ok {
		if defaultHost != "" {
			ep = defaultHost
		} else {
			return endpoint{}, fmt.Errorf("environment variable %s is not set", perforatorEndpointEnv)
		}
	}

	if level, ok := os.LookupEnv(perforatorSecurityLevelEnv); ok {
		secure, err := strconv.ParseBool(level)
		if err != nil {
			return endpoint{}, fmt.Errorf(
				"failed to parse %s: expected bool, found %s", perforatorSecurityLevelEnv, level,
			)
		}
		return parseEndpoint(ep, &secure)
	} else {
		return parseEndpoint(ep, nil)
	}
}
