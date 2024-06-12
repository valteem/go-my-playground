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

// The following is set to benchmark two FNV hashing algorithms, regular and `fast`

const (
	offset64 = 14695981039346656037
	prime64  = 1099511628211
)

var (
	emptyLabelSignature = hashNew()
)

// hashNew initializes a new fnv64a hash value.
func hashNew() uint64 {
	return offset64
}

// hashAdd adds a string to a fnv64a hash value, returning the updated hash.
func hashAdd(h uint64, s string) uint64 {
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= prime64
	}
	return h
}

// hashAddByte adds a byte to a fnv64a hash value, returning the updated hash.
func hashAddByte(h uint64, b byte) uint64 {
	h ^= uint64(b)
	h *= prime64
	return h
}

func labelsToSignature(labels map[string]string) uint64 {
	if len(labels) == 0 {
		return emptyLabelSignature
	}

	labelNames := make([]string, 0, len(labels)) // allocation
	for labelName := range labels {
		labelNames = append(labelNames, labelName)
	}
	sort.Strings(labelNames)

	sum := hashNew()
	for _, labelName := range labelNames {
		sum = hashAdd(sum, labelName)
		sum = hashAddByte(sum, separatorByte)
		sum = hashAdd(sum, labels[labelName])
		sum = hashAddByte(sum, separatorByte)
	}
	return sum
}

func labelsToSignatureFast(m map[string]string) uint64 {
	if len(m) == 0 {
		return emptyLabelSignature
	}

	var result uint64
	for labelName, labelValue := range m {
		sum := hashNew()
		sum = hashAdd(sum, labelName)
		sum = hashAddByte(sum, separatorByte)
		sum = hashAdd(sum, labelValue)
		result ^= sum
	}
	return result
}

// go test fnv_test.go -gcflags -'-m=3 -l'
func BenchmarkLabelsToSignature(b *testing.B) {
	m := map[string]string{
		"label1":               "value1",
		"label2":               "value2",
		"label3":               "value3",
		"some random label":    "some random value",
		"another random label": "another random vaule",
		"other label stuff":    "other value stuff",
		"more somple labels":   "more simple values",
	}

	b.Run("Regular labels-to-signature", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = labelsToSignature(m)
		}
	})

	b.Run("Fast labels-to-signature", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = labelsToSignatureFast(m)
		}
	})
}
