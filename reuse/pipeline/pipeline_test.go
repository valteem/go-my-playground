package pipeline

import (
	"testing"
)

func collectParallel(n int) []foodRecord {

	c := ServeFoodParallel(n)
	rec := []foodRecord{}

	for output := range c {
		rec = append(rec, output)
	}

	return rec

}

func TestServeFoodParallel(t *testing.T) {

	tests := []struct {
		n int
	}{
		{10}, {100}, {1000}, {10000}, {100000}, {1000000},
	}

	for _, tc := range tests {
		r := collectParallel(tc.n)
		if len(r) != tc.n {
			t.Errorf("number of dishes: get %d, expect  %d", len(r), tc.n)
		}

	}

}

func collectBounded(n int, numCookers int) []foodRecord {

	c := ServeFoodBounded(n, numCookers)
	rec := []foodRecord{}

	for output := range c {
		rec = append(rec, output)
	}

	return rec

}

func TestServeFoodBounded(t *testing.T) {

	tests := []struct {
		n         int
		numCooker int
	}{
		{10, 1}, {10, 10}, {100, 10}, {1000, 10}, {10000, 10}, {100000, 10},
	}

	for _, tc := range tests {
		r := collectBounded(tc.n, tc.numCooker)
		if len(r) != tc.n {
			t.Errorf("number of dishes: get %d, expect  %d", len(r), tc.n)
		}

	}

}

func BenchmarkPipeline(b *testing.B) {

	b.Run("parallel, 100 dishes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			collectParallel(100)
		}
	})

	b.Run("bounded, 100 dishes, 1 cooker", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			collectBounded(100, 1)
		}
	})

	b.Run("bounded, 100 dishes, 5 cookers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			collectBounded(100, 5)
		}
	})

	b.Run("bounded, 100 dishes, 10 cookers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			collectBounded(100, 10)
		}
	})

	b.Run("bounded, 100 dishes, 30 cookers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			collectBounded(100, 30)
		}
	})

	b.Run("bounded, 100 dishes, 50 cookers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			collectBounded(100, 50)
		}
	})

}

func TestMD5All(t *testing.T) {

	var (
		inputSize   = 100
		inputLength = 10
		numWorkers  = 10
	)

	input := randomByteArray(inputSize, inputLength)

	sum := MD5AllParallel(input)
	if sum != inputSize {
		t.Errorf("parallel: get %d, expect %d", sum, inputSize)
	}

	sum = MD5AllBounded(input, numWorkers)
	if sum != inputSize {
		t.Errorf("bounded: get %d, expect %d", sum, inputSize)
	}

}

func BenchmarkMD5Pipeline(b *testing.B) {

	input := randomByteArray(10, 10)

	b.Run("serial", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MD5AllSerial(input)
		}
	})

	b.Run("parallel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MD5AllParallel(input)
		}
	})

	b.Run("bounded", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MD5AllBounded(input, 2)
		}
	})

}

func BenchmarkSquaresPipeline(b *testing.B) {

	b.Run("serial", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SquaresSerial(1000000)
		}
	})

	b.Run("parallel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SquaresParallel(1000000)
		}
	})

	b.Run("bounded", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SquaresBounded(1000000, 8)
		}
	})

}

func TestSplitMerge(t *testing.T) {

	var numMsg = 10
	c := make(chan int)

	split := Split(c, 2)

	merge := Merge(split...)

	go func() {
		for i := range numMsg {
			c <- i
		}
		close(c)
	}()

	count, sum := 0, 0

	for v := range merge {
		sum += v
		count++
	}

	if count != numMsg || sum != (numMsg)*(numMsg-1)/2 {
		t.Errorf("get sum/count %d/%d, expect %d/%d", sum, count, (numMsg)*(numMsg-1)/2, numMsg)
	}

}
