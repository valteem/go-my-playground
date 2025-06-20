package main

import (
	"strings"
	"testing"

	"golang.org/x/text/language"

	"github.com/pbnjay/memory"
)

func TestTotalmemory(t *testing.T) {

	totalMemory, freeMemory := memory.TotalMemory(), memory.FreeMemory()

	if totalMemory < freeMemory {
		t.Errorf("total memory %d, free memory %d", totalMemory, freeMemory)
	}

	totalMemoryMiB, freeMemoryMiB := MemorySizeFormatted(totalMemory, MiB, language.English), MemorySizeFormatted(freeMemory, MiB, language.English)
	totalMemoryDecimal, freeMemoryDecimal :=
		strings.TrimSuffix(totalMemoryMiB, " MiB"), strings.TrimSuffix(freeMemoryMiB, " MiB")
	if len(totalMemoryDecimal) == 0 || len(freeMemoryDecimal) == 0 {
		t.Errorf("decimal memory representation cannot be empty")
	}

}

func TestMemoryUnits(t *testing.T) {

	tests := []struct {
		inputSize     uint64
		inputUnits    MemoryUnits
		inputLanguage language.Tag
		output        string
	}{
		{123456789, KiB, language.English, "120,563 KiB"},
		{123456789, MiB, language.Serbian, "118 MiB"},
	}

	for _, tc := range tests {
		if output := MemorySizeFormatted(tc.inputSize, tc.inputUnits, tc.inputLanguage); output != tc.output {
			t.Errorf("get %q, expect %q", output, tc.output)
		}
	}
}
