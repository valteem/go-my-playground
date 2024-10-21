package reuse

import (
	"io"
)

type MultiWriter struct {
	writers []io.Writer
}

func (mw MultiWriter) Write(b []byte) error {
	for _, writer := range mw.writers {
		_, err := writer.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (mw *MultiWriter) AddWriter(w io.Writer) {
	mw.writers = append(mw.writers, w)
}
