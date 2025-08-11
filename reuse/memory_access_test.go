package reuse_test

import (
	"math/rand"
	"testing"
)

// go test -test.bench BenchmarkMemoryAccessReadOnly -test.run=neverforever . -v -gcflags='-N -l' -benchmem
func BenchmarkMemoryAccessReadOnly(b *testing.B) {

	b.Run("pre-allocate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			m := make([]int32, 0, 1<<25+1)
			for range 1 << 25 {
				m = append(m, rand.Int31())
			}
			b.StartTimer()
			for _, v := range m {
				v += 1
			}
		}
	})

	b.Run("set-direct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			m := make([]int32, 1<<25)
			for j := range 1 << 25 {
				m[j] = rand.Int31()
			}
			b.StartTimer()
			for _, v := range m {
				v += 1
			}
		}
	})

	b.Run("resize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			var m []int32
			for range 1 << 25 {
				m = append(m, rand.Int31())
			}
			b.StartTimer()
			for _, v := range m {
				v += 1
			}
		}
	})

}

// go test -test.bench BenchmarkMemoryAccessReadWrite -test.run=neverforever . -v -gcflags='-N -l' -benchmem
func BenchmarkMemoryAccessReadWrite(b *testing.B) {

	b.Run("pre-allocate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			// m := make([]int32, 0, 38500352) // same as "resize result"
			m := make([]int32, 0, 2<<25)
			for range 1 << 25 {
				m = append(m, rand.Int31())
			}
			b.StartTimer()
			for j := range m {
				m[j] += 1
			}
		}
	})

	b.Run("set-direct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			m := make([]int32, 1<<25)
			for j := range 1 << 25 {
				m[j] = rand.Int31()
			}
			b.StartTimer()
			for j := range m {
				m[j] += 1
			}
		}
	})

	b.Run("resize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			var m []int32
			for range 1 << 25 {
				m = append(m, rand.Int31())
			}
			b.StartTimer()
			for j := range m {
				m[j] += 1
			}
		}
	})

	b.Run("resize-modify", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			var m []int32
			for range 1 << 25 {
				m = append(m, rand.Int31())
			}
			m[0] += 1 // reset memory controller buffers
			m[0] -= 1
			b.StartTimer()
			for j := range m {
				m[j] += 1
			}
		}
	})

}

/*
Possible Low-Level Explanations:
Write buffer priming: The initial operations "prime" the CPU's write buffers in an optimal way
DRAM row buffer optimization: The memory access pattern leaves DRAM row buffers in a favorable state
Cache controller state: The cache replacement algorithms or prefetchers are in an optimal state
TLB entry organization: The translation lookaside buffer has optimal entries for the memory region
This is Hardware-Specific Behavior
This kind of performance difference is highly dependent on:

Specific CPU architecture
Memory controller design
Cache hierarchy implementation
Memory channel configuration (single vs dual-channel)
*/

func augmentMemoryAllocations(size int) {
	m := make([]int32, 0, 2<<size)
	for range 2 << size {
		m = append(m, rand.Int31())
	}
	m[0] += 1
	m[2<<size-1] += 1
}

func BenchmarkMemoryAccessReadWriteAugmented(b *testing.B) {

	b.Run("pre-allocate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			// m := make([]int32, 0, 38500352) // same as "resize result"
			m := make([]int32, 0, 2<<25)
			for range 1 << 25 {
				m = append(m, rand.Int31())
			}
			augmentMemoryAllocations(27)
			b.StartTimer()
			for j := range m {
				m[j] += 1
			}
		}
	})

	b.Run("set-direct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			m := make([]int32, 1<<25)
			for j := range 1 << 25 {
				m[j] = rand.Int31()
			}
			augmentMemoryAllocations(27)
			b.StartTimer()
			for j := range m {
				m[j] += 1
			}
		}
	})

	b.Run("resize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			var m []int32
			for range 1 << 25 {
				m = append(m, rand.Int31())
			}
			augmentMemoryAllocations(27)
			b.StartTimer()
			for j := range m {
				m[j] += 1
			}
		}
	})

	b.Run("resize-modify", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			var m []int32
			for range 1 << 25 {
				m = append(m, rand.Int31())
			}
			m[0] += 1 // reset memory controller buffers
			m[0] -= 1
			augmentMemoryAllocations(27)
			b.StartTimer()
			for j := range m {
				m[j] += 1
			}
		}
	})

}
func BenchmarkMemoryPopulate(b *testing.B) {

	b.Run("pre-allocate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := make([]int32, 0, 1<<25+1) // (1<<25) produce 5 allocs/op
			for range 1 << 25 {
				//lint:ignore SA4010 for test purposes only
				m = append(m, rand.Int31())
			}
		}
	})

	b.Run("set-direct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := make([]int32, 1<<25)
			for j := range 1 << 25 {
				m[j] = rand.Int31()
			}
		}
	})

	b.Run("resize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var m []int32
			for range 1 << 25 {
				//lint:ignore SA4010 for test purposes only
				m = append(m, rand.Int31())
			}
		}
	})
}
