package pipegen

import (
	"sync"
)

type cmd func(in, out chan any)

func RunPipeline(cmds ...cmd) {

	var wg sync.WaitGroup

	wg.Add(len(cmds))

	ins := make([]chan any, 1) // need first element to check it for non-nilness
	outs := make([]chan any, 0)

	for i, cm := range cmds {
		if ins[i] == nil {
			ins[i] = make(chan any) // first cmd generates out without any input
		}
		outs = append(outs, make(chan any))
		go func() {
			j := i // probably not needed after introduction of new loop capture behaviour
			cm(ins[j], outs[j])
			close(outs[j])
			wg.Done()
		}()
		ins = append(ins, outs[i])
	}

	wg.Wait()

}
