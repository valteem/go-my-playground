// https://github.com/prometheus/common/blob/6b9921f9eba2cd74f2caca0d713bb0a6eb7ef1b9/model/labelset_string_go120.go#L24

// Benchmark two versions of LabelSet.String()

package reuse_test

import (
	"bytes"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
	"testing"
)

type LabelName string
type LabelValue string
type LabelSet map[LabelName]LabelValue

func (l LabelSet) StringNew() string {
	var lna [32]string // On stack to avoid memory allocation for sorting names.
	labelNames := lna[:0]
	for name := range l {
		labelNames = append(labelNames, string(name))
	}
	slices.Sort(labelNames)
	var bytea [1024]byte // On stack to avoid memory allocation while building the output.
	b := bytes.NewBuffer(bytea[:0])
	b.WriteByte('{')
	for i, name := range labelNames {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(name)
		b.WriteByte('=')
		b.Write(strconv.AppendQuote(b.AvailableBuffer(), string(l[LabelName(name)]))) // buffer.AvailableBuffer() added in go 1.21.0
	}
	b.WriteByte('}')
	return b.String()
}

func (l LabelSet) StringOld() string {
	labelNames := make([]string, 0, len(l))
	for name := range l {
		labelNames = append(labelNames, string(name))
	}
	sort.Strings(labelNames)
	lstrs := make([]string, 0, len(l))
	for _, name := range labelNames {
		lstrs = append(lstrs, fmt.Sprintf("%s=%q", name, l[LabelName(name)]))
	}
	return fmt.Sprintf("{%s}", strings.Join(lstrs, ", "))
}

func PopulateLabelSet(l int) LabelSet {
	m := LabelSet{}
	for i := 0; i < l; i++ {
		name, value := LabelName(fmt.Sprintf("name%08d", i)), LabelValue(fmt.Sprintf("value%07d", i))
		m[name] = value
	}
	return m
}

func BenchmarkString(b *testing.B) {
	ls := PopulateLabelSet(1024 * 1024)

	b.Run("StringNew", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ls.StringNew()
		}
	})

	b.Run("StringOld", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ls.StringOld()
		}
	})

}
