package main

import (
	"math"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type MemoryUnits int

const (
	KiB MemoryUnits = iota
	MiB
	TiB

	divBytes uint64 = 1
	divKiB   uint64 = 1024
	divMiB   uint64 = 1024 * 1024
	divTiB   uint64 = 1024 * 1024 * 1024

	unitBytes = "Bytes"
	unitKiB   = "KiB"
	unitMiB   = "MiB"
	unitTiB   = "TiB"
)

func MemorySizeFormatted(size uint64, format MemoryUnits, language language.Tag) string {

	var divider uint64
	var unit string

	switch format {
	case KiB:
		divider = divKiB
		unit = unitKiB
	case MiB:
		divider = divMiB
		unit = unitMiB
	case TiB:
		divider = divTiB
		unit = unitTiB
	default:
		divider = divBytes
		unit = unitBytes
	}

	printer := message.NewPrinter(language)

	return printer.Sprintf("%d %s", uint64(math.Round(float64(size)/float64(divider))), unit)

}
