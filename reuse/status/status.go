package status

type State uint32

const (
	NotReady State = iota
	Ready
	Stopping
	Idle
)
