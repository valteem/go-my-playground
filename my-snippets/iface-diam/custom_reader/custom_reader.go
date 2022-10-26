package custom_reader

type CustomReader interface {
	Open(message_open string)
	Close(message_close string)
}