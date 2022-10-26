package bufiom

import "fmt"

type SomeReader interface{
	Read() int // number of bytes flushed
}

type BufReader struct {
	Buf []byte
	Size int
}

func NewBufReader(size int) *BufReader {
	b := BufReader{Buf: make([]byte, size), Size: size}
	return &b
}

func NewBufReaderVal(size int) BufReader {
	b := BufReader{Buf: make([]byte, size), Size: size}
	return b
}

func (r BufReader) Read() int {
	s := len(r.Buf)
	return s
}

func (r *BufReader) Overwrite(s string) {
	r.Buf = nil
	b := []byte(s)
	r.Buf = append(r.Buf, b...)
	r.Size = len(b)
}

type OtherReader struct {
	Buf []byte
	Size int
}

func NewOtherReader(size int) *OtherReader {
	o := OtherReader{Buf: make([]byte, size), Size: size}
	return &o
}

func NewOtherReaderVal(size int) OtherReader {
	o := OtherReader{Buf: make([]byte, size), Size: size}
	return o
}

func (r OtherReader) Read() int {
	s := len(r.Buf)
	return s
}

func (r *OtherReader) Overwrite(s string) {
	r.Buf = nil
	b := []byte(s)
	r.Buf = append(r.Buf, b...)
	r.Size = len(b)
}

func AssertReaderType(descr string, r SomeReader) {
	fmt.Println(descr, r)
	brr, okrr := r.(*BufReader)
	brv, okrv := r.(BufReader)
	fmt.Println("Type assertion - *BufReader", brr, okrr)
	fmt.Println("Type assertion - BufReader", brv, okrv)
	bor, okor := r.(*OtherReader)
	bov, okov := r.(OtherReader)
	fmt.Println("Type assertion - *OtherReader", bor, okor)
	fmt.Println("Type assertion - OtherReader", bov, okov)
}