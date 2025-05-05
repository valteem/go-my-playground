package pipeline

import (
	"crypto/md5"
	"math/rand"
	"strings"
	"sync"
)

var (
	food   = []string{"apples", "onions", "cucumbers", "pears", "cherries", "potatoes", "tomatoes", "pizzas", "cakes", "baguets", "gingers", "garlics", "carrots"}
	tastes = []string{"sweet", "salted", "sour", "smoked"}
	runes  = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}[]")
)

type foodRecord struct {
	food string
	idx  int
}

func ServeFoodParallel(n int) <-chan foodRecord {
	c := make(chan foodRecord)
	var wg sync.WaitGroup
	go func() {
		count := 0
		for count < n {
			raw := food[rand.Intn(len(food)-1)]
			wg.Add(1)
			count++
			idx := count
			go func() {
				taste := tastes[rand.Intn(len(tastes))]
				sb := strings.Builder{}
				sb.WriteString(taste)
				sb.WriteString(" ")
				sb.WriteString(raw)
				c <- foodRecord{sb.String(), idx}
				wg.Done()
			}()
		}
		go func() {
			wg.Wait()
			close(c)
		}()
	}()
	return c
}

func rawFood(n int) <-chan foodRecord {
	c := make(chan foodRecord)
	var wg sync.WaitGroup
	go func() {
		count := 0
		for count < n {
			raw := food[rand.Intn(len(food)-1)]
			wg.Add(1)
			count++
			idx := count
			go func() {
				c <- foodRecord{raw, idx}
				wg.Done()
			}()
		}
		go func() {
			wg.Wait()
			close(c)
		}()
	}()
	return c
}

func cooker(input <-chan foodRecord, output chan<- foodRecord) {
	sb := strings.Builder{}
	for inp := range input {
		taste := tastes[rand.Intn(len(tastes))]
		sb.WriteString(taste)
		sb.WriteString(" ")
		sb.WriteString(inp.food)
		output <- foodRecord{sb.String(), inp.idx}
		sb.Reset()
	}
}

func ServeFoodBounded(n int, numCookers int) <-chan foodRecord {
	input := rawFood(n)
	output := make(chan foodRecord)
	var wg sync.WaitGroup
	wg.Add(numCookers)
	for i := 0; i < numCookers; i++ {
		go func() {
			cooker(input, output)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

func Merge(inputs ...<-chan string) <-chan string {

	var wg sync.WaitGroup
	output := make(chan string)

	out := func(input <-chan string) {
		for s := range input {
			output <- s
		}
		wg.Done()
	}

	wg.Add(len(inputs))
	for _, c := range inputs {
		out(c)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output

}

func randStr(length int) []byte {
	s := make([]byte, length)
	for i := range s {
		s[i] = byte(runes[rand.Intn(len(runes))])
	}
	return []byte(s)
}

func randomByteArray(n, length int) [][]byte {
	output := make([][]byte, n)
	for i := range output {
		output[i] = randStr(length)
	}
	return output
}

func MD5SumParallel(input [][]byte) <-chan [md5.Size]byte {
	c := make(chan [md5.Size]byte)
	var wg sync.WaitGroup
	go func() {
		idx := 0
		for idx < len(input) {
			wg.Add(1)
			jdx := idx
			go func() {
				c <- md5.Sum(input[jdx])
				wg.Done()
			}()
			idx++
		}
		go func() {
			wg.Wait()
			close(c)
		}()
	}()
	return c

}

func MD5AllParallel(input [][]byte) int {

	c := MD5SumParallel(input)

	count := 0
	for range c {
		count++
	}

	return count

}

func byteToMD5(input <-chan []byte, output chan<- [md5.Size]byte, count chan int) {
	for v := range input {
		output <- md5.Sum(v)
		count <- 1
	}
}

func serveBytes(input [][]byte) <-chan []byte {
	c := make(chan []byte)
	var wg sync.WaitGroup
	go func() {
		idx := 0
		for idx < len(input) {
			wg.Add(1)
			jdx := idx
			go func() {
				c <- input[jdx]
				wg.Done()
			}()
			idx++
		}
		go func() {
			wg.Wait()
			close(c)
		}()
	}()
	return c
}

func MD5SumBounded(input [][]byte) <-chan [md5.Size]byte {
	c := make(chan [md5.Size]byte)
	var wg sync.WaitGroup
	idx := 0
	for idx < len(input) {
		jdx := idx
		wg.Add(1)
		go func() {
			c <- md5.Sum(input[jdx])
			wg.Done()
		}()
		idx++
	}
	go func() {
		wg.Wait()
		close(c)
	}()

	return c

}

func MD5CollectBounded(input [][]byte, numWorkers int) <-chan [md5.Size]byte {
	inp := serveBytes(input)
	rec := make(chan [md5.Size]byte)
	count := make(chan int)

	for range numWorkers {
		go func() {
			byteToMD5(inp, rec, count)
		}()
	}

	go func() {
		counter := 0
		for v := range count {
			counter += v
			if counter == len(input) {
				close(rec)
				close(count)
				return
			}
		}
	}()

	return rec

}

func MD5AllBounded(input [][]byte, numWorkers int) int {

	c := MD5CollectBounded(input, numWorkers)

	sum := 0
	for range c {
		sum++
	}

	return sum
}

func MD5AllSerial(input [][]byte) int {

	sum := 0
	for _, v := range input {
		_ = md5.Sum(v)
		sum++
	}

	return sum

}

func SquaresSerial(n int) int {
	output := 0
	for v := range n {
		output += v * v
	}
	return output
}

func SquaresParallel(n int) int {
	c := make(chan int)
	for v := range n {
		go func() {
			c <- v * v
		}()
	}
	count, output := 0, 0
	for v := range c {
		output += v
		count++
		if count == n {
			break
		}
	}
	return output
}

func SquaresBounded(n int, workers int) int {
	c, r := make(chan int), make(chan int)
	for v := range n {
		go func() {
			c <- v
		}()
	}
	for range workers {
		go func(inp, out chan int) {
			for v := range inp {
				out <- v * v
			}
		}(c, r)

	}
	output, count := 0, 0
	for v := range r {
		output += v
		count++
		if count == n {
			break
		}
	}
	return output
}
