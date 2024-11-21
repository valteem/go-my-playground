package reuse

import (
	"encoding/json"
)

// https://stackoverflow.com/a/76923923
func ParseAnyJSON(data string) (any, error) {
	var object any
	err := json.Unmarshal([]byte(data), &object)
	if err != nil {
		return nil, err
	}
	return object, nil
}
