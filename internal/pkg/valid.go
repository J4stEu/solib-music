package pkg

import (
	"net"
	"net/url"
)

func IsValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func IsValidIP(toTest string) bool {
	if toTest != "localhost" {
		if net.ParseIP(toTest) == nil {
			return false
		}
	}
	return true
}
