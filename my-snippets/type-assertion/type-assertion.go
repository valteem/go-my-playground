package main

import "fmt"

func main() {
	var x interface{} = 7
	i := x.(int)
	// Checking 'escape to heap' with go build -gcflags="-m"
	fmt.Println(i) // i escapes to heap this way as it is passsed to Println function
	j := x.(int)
	j += 1 // i does not escape to heap as it is not passed anywhere outside main() function

	if ai, ok := x.(int); ok {
		fmt.Println("this is really int ",ai, ok)
	}

	var s interface{} = "str1"
	if as, ok := s.(string); ok {
		fmt.Println("this is really string ", as, ok)
	}

	var b any = []byte("word")
    fmt.Println(b)
//	fmt.Println(b[1]) // This does not work, cannot index variable type any
	if ab, ok := b.(byte); ok {
		fmt.Println("this is really byte", ab, ok)	
	} else {
		fmt.Println("This is something different", ab, ok)
	}
	b1 := []byte("word")
	fmt.Println(b1[0], rune(b1[0]), string(b1[0]), string(rune(b1[0]))) // This works

	var b2 []byte // This way the slice has both length and capacity zero
	copy(b2, []byte("word")) // this is why copy function copies nothing
	fmt.Println(len(b2), cap(b2), b2)
	var ba2[]byte
	ba2 = append(ba2, "word"...)
	fmt.Println(len(ba2), cap(ba2), ba2, string(ba2))

	b3 := make([]byte, 3) // This way new slice has length and capacity 3
	copy(b3, "word") // this is why copy function copies 3 symbols
	fmt.Println(len(b3), cap(b3), b3, string(b3))

	b4 := make([]byte, 3, 5)
	copy(b4, "word")
	fmt.Println(len(b4), cap(b4), b4, string(b4))

	var b5[]byte // byte slice grows capacity to 8 after first append
	fmt.Println(0, len(b5), cap(b5))
	for i := 1; i <= 5; i++ {
		b5 = append(b5, []byte(fmt.Sprint(i))...)
		fmt.Println(i, len(b5), cap(b5), b5, string(b5))
	}

	var i5[]int // int slice increases capacity by 1 during first 4 appends, then jumps to 8 at fifth append
	fmt.Println(0, len(i5), cap(i5))
	for i := 1; i <= 5; i++ {
		i5 = append(i5, i)
		fmt.Println(i, len(i5), cap(i5), i5)
	}

	var s5 []string // string slice shows the same as byte slice
	fmt.Println(0, len(s5), cap(s5), s5)
	for i := 1; i <= 5; i++ {
		s5 = append(s5, fmt.Sprint(i))
		fmt.Println(i, len(s5), cap(s5), s5)
	}

	type combo struct {
		txt string
		num int
		num2 int
	}

	var c5 []combo // custom struct slice jumps to capacity 4 at third append, to 8 at fifth append, and to 16 at ninth
	fmt.Println(0, len(c5), cap(c5), c5)
	for i := 1; i <= 10; i++ {
		c5 = append(c5, combo{fmt.Sprint(i),i, 2*i})
		fmt.Println(i, len(c5), cap(c5), c5)
	}	

}