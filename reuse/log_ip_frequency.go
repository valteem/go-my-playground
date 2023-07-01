package reuse

import (
	"net"
	"regexp"
)

const (
	pat = `\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`
)

func validateIPAddress(s string) bool { return net.ParseIP(s) != nil }

func ParseLogForIP(l string) map[string]int {

	re, _ := regexp.Compile(pat) // defined in ipv4_validate.go

	matches := re.FindAllString(l, -1)

	count := make(map[string]int)

	for _, m := range matches {
		if validateIPAddress(m) {
			count[m]++
		}
	}

	return count
}