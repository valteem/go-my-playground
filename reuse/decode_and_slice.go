package reuse

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// parses base64-encoded string, returns slice of values initially separated with a delimiter
func DecodeAndSlice(input string, delimiter string) ([]string, error) {

	s, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		var dummy []string
		return dummy, err
	}

	decoded := string(s)

	if len(decoded) < len(delimiter) {
		var dummy []string
		return dummy, fmt.Errorf("delimiter cannot be longer than decoded input")
	}

	return strings.Split(decoded, delimiter), nil

}
