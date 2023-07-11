package reuse

func CopyUnderlyingBuffer(s string) string {

	buffer := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		buffer[i] = s[i]
	}
	return string(buffer)
	
}