package reuse

type Mux struct {
	Path    []string
	Handler []func()
}

var newMux Mux
var ExportedMux = &newMux

func ExportMux(path []string, handler []func()) *Mux {

	ExportedMux.Path = path
	ExportedMux.Handler = handler

	return ExportedMux

}
