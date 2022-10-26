package main

import (
	"lang.rev/iface-diam/custom_reader"
	"lang.rev/iface-diam/custom_writer"
)

type CustomReaderWriter interface {
	custom_reader.CustomReader
	custom_writer.CustomWriter
}