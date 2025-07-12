package app

import (
	"github.com/valteem/reuse/initorder/explicit/declare"
	_ "github.com/valteem/reuse/initorder/explicit/use"
)

type app struct {
	url string
}

func NewApp() *app {
	return &app{url: declare.URL}
}
