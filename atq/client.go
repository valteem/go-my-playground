package atq

type OptionType int

const (
	MaxRetryOpt OptionType = iota
	QueueOpt
	TimeoutOpt
)

type Option interface {
	String() string
	Type() Option
	Value() any // value used to create the option
}