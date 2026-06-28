package singleton

var globalPipe *MultiPipe

func SetGlobalPipe(mp *MultiPipe) {
	globalPipe = mp
}
func Validate() error {
	if globalPipe == nil {
		panic("globalPipe must be set before calling Validate")
	}
	return globalPipe.Validate()
}

func Read() string {
	if globalPipe == nil {
		panic("globalPipe must be set before calling Read")
	}
	return globalPipe.Read()
}
