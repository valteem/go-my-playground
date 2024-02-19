package base

// Internal representation of a task, with additional metadata
type TaskMessage struct {
	Type string
	Payload []byte
	ID string
	Queue string
	Retry int
	Retried int
	ErrorMsg string
	LastFailedAt int64 // Unix time format, zero if no failures
	Timeout int64
	Deadline int64
	UniqueKey string // redis key for uniqueness lock, empty if no lock was used
	GroupKey string // TaskInfo.Group ? empty if no aggregation is planned for the task
	Retention int64 // how long (in seconds) retain the task after completion
	CompletedAt int64
}