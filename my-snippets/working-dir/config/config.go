package config

import (
	"os"
)

var (
	ConfigPath string
)

func init() {
	ConfigPath, _ = os.Getwd()
}
