package reuse

import (
	"runtime"
)
type Platform struct {
	OS string
}

type Compile struct {
	Platform
}

type Builder struct {
	Compile
}

func NewPlatform() *Platform {
	return &Platform{OS:runtime.GOOS}
}

func NewCompile() *Compile{
	return &Compile{*NewPlatform()}
}

func NewBuilder() *Builder { 
	return &Builder{*NewCompile()}
}