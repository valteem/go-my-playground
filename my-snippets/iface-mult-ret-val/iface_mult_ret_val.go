package main

import "fmt"

type multRet interface {
	Output() (int, int)
}

type Example struct {
	param int
}

func (e Example) Output() (int, int) {
	return e.param*e.param, e.param*e.param*e.param
}


func show(mr multRet) {
	fmt.Println(mr.Output())
}

func main(){

	e := Example{param: 2}
	show(e)

}