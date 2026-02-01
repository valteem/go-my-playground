package pipegen

import (
	"strings"

	"testing"
)

const (
	strNum = 100
)

var (
	inpStr = []string{"it's a pity ", "little kitty ", "lives in big sity"}
	outStr = strings.Join(inpStr, "")
)

func buildAddStr(str string) cmd {
	return func(in, out chan any) {
		for inp := range in {
			inStr := inp.(string)
			out <- inStr + str
		}
	}
}

func buildSrc(i int) cmd {
	return func(in, out chan any) {
		var s string
		for range i {
			out <- s
		}
	}
}

func buildRecv(r *[]string) cmd {
	return func(in, out chan any) {
		for inp := range in {
			*r = append(*r, inp.(string))
		}
	}
}

func initPipeline(r *[]string) []cmd {
	output := []cmd{}
	var c cmd
	c = buildSrc(strNum)
	output = append(output, c)
	for _, s := range inpStr {
		c = buildAddStr(s)
		output = append(output, c)
	}
	c = buildRecv(r)
	output = append(output, c)
	return output
}

func TestPipeStr(t *testing.T) {

	s := []string{}
	r := &s

	cmds := initPipeline(r)

	RunPipeline(cmds...)

	for i, v := range s {
		if v != outStr {
			t.Errorf("%d: get %q, expect %q", i, v, outStr)
		}
	}
}
