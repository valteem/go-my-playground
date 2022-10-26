package custom_writer

type CustomWriter interface {
	Open(message_open string)
	Close(message_close string)
}