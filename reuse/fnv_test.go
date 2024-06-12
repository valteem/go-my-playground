package reuse_test

import (
	"bytes"
	"hash"
	"hash/fnv"
	"sort"
	"testing"
)

const (
	separatorByte = 255
	base64        = 14695981039346656037
)

func MapToFNV(m map[string]string) uint64 {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var b bytes.Buffer
	var h hash.Hash64 = fnv.New64a()
	for _, key := range keys {
		b.WriteString(key)
		b.WriteByte(separatorByte)
		b.WriteString(m[key])
		b.WriteByte(separatorByte)
		h.Write(b.Bytes())
		b.Reset()
	}
	return h.Sum64()
}

func TestMapToFNV(t *testing.T) {
	tests := []struct {
		input  map[string]string
		output uint64
	}{
		{
			input:  map[string]string{},
			output: base64,
		},
		{
			input:  map[string]string{"name": "garland, briggs", "fear": "love is not enough"},
			output: 5799056148416392346,
		},
	}
	for _, tc := range tests {
		if output := MapToFNV(tc.input); output != tc.output {
			t.Errorf("%v MapToFNV(): get %d, expect %d", tc.input, output, tc.output)
		}
	}
}
