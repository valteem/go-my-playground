package channelsm

import "sync"

func MergeChannels(chin ...<-chan string) <-chan string {

	var wg sync.WaitGroup
	wg.Add(len(chin))
	chout := make(chan string)
	for _, c := range chin {
		go func (c <-chan string) {
			for v := range c {
				chout <- v	
    		}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(chout)
	}()

	return chout

}