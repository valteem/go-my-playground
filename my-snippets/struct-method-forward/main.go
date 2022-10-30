package main

import (
	"fmt"
)

type Pusher struct {
	content string
}

func (p *Pusher) Init(s string) {
	p.content = s
}

func (p Pusher) Push() {
	fmt.Println("Pushing", p.content)
}

type Puller struct {
	content string
}

func (p *Puller) Init(s string) {
	p.content = s
}

func (p Puller) Pull() {
	fmt.Println("Pulling", p.content)
}

type PusherPuller struct {
	// *Pusher
	// *Puller
	Pusher
	Puller
}

func main() {

	// pp := PusherPuller{Pusher: new(Pusher), Puller: new(Puller)}
	push := new(Pusher)
	pull := new(Puller)
	pp := PusherPuller{Pusher: *push, Puller: *pull}

	pp.Pusher.Init("Push !!!")
	pp.Puller.Init("Pull !!!")

	pp.Push()
	pp.Pull()

}