// https://stackoverflow.com/questions/48597289/printf-security-golang?rq=1

package main
import (
	"fmt"
	"strconv"
)

func main() {

//  \x08 (also \b) is an escape code for Backspace ASCII character
	inputString := "An input string with some\x08\b\x08\bhidden characters"
	fmt.Printf("%s\n", inputString) // %s is a verb for string values
	fmt.Printf("%q\n", inputString) // %q is 'a double-quoted string safely escaped with Go syntax'
	fmt.Printf("%s\n",strconv.Quote(inputString))
}