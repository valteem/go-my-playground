package main

import (
	"bytes"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
	"testing"
)

const (
	amin, amax = 1, 1024 * 1024
)

type LabelSet map[string]string

func populateSet(nmin, nmax int) LabelSet {
	m := LabelSet{}
	for i := nmin; i <= nmax; i++ {
		name, value := fmt.Sprintf("name%8d", i), fmt.Sprintf("value%8d", i)
		m[name] = value
	}
	return m
}

func stringNamesNew(ls LabelSet) string {
	var lna [32]string
	labelNames := lna[:0]
	for name := range ls {
		labelNames = append(labelNames, string(name))
	}
	slices.Sort(labelNames)
	s := strings.Join(labelNames, ",")
	return s
}

func stringNamesOld(ls LabelSet) string {
	labelNames := make([]string, 0, len(ls)) // escapes to heap
	for name := range ls {
		labelNames = append(labelNames, string(name))
	}
	sort.Strings(labelNames)
	s := strings.Join(labelNames, ",")

	return s
}

func stringNamesValuesNew(ls LabelSet) string {
	var lna [32]string
	labelNames := lna[:0]
	for name := range ls {
		labelNames = append(labelNames, string(name))
	}
	slices.Sort(labelNames)
	// %name%=%value%
	var bytea [1024]byte // On stack to avoid memory allocation while building the output.
	b := bytes.NewBuffer(bytea[:0])
	b.WriteByte('{')
	for i, name := range labelNames {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(name)
		b.WriteByte('=')
		b.Write(strconv.AppendQuote(b.AvailableBuffer(), string(ls[string(name)]))) // buffer.AvailableBuffer() added in go 1.21.0
	}
	b.WriteByte('}')
	return b.String()
}

func stringNamesValuesOld(ls LabelSet) string {
	labelNames := make([]string, 0, len(ls)) // escapes to heap
	for name := range ls {
		labelNames = append(labelNames, string(name))
	}
	sort.Strings(labelNames)
	// %name%=%value%
	lstrs := make([]string, 0, len(ls))
	for _, name := range labelNames {
		lstrs = append(lstrs, fmt.Sprintf("%s=%q", name, ls[string(name)]))
	}
	return fmt.Sprintf("{%s}", strings.Join(lstrs, ", "))
}
func BenchmarkAllocations(b *testing.B) {
	nv := populateSet(amin, amax)
	b.Run("names - new", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = stringNamesNew(nv)
		}
	})
	b.Run("names - old", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = stringNamesOld(nv)
		}
	})
	b.Run("names and values - new", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = stringNamesValuesNew(nv)
		}
	})
	b.Run("names and values - old", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = stringNamesValuesOld(nv)
		}
	})
}
