// https://leighmcculloch.com/posts/go-finding-escapes-to-heap/
// https://github.com/golang/vscode-go/wiki/settings

// gopls:
//     ui.codelenses
//         gc_details
//    ui.diagnostic.annotations
//        escape

// go test var_decl_test.go -gcflags "-m -m"  -bench . -benchmem -memprofile mem.out

// both BeforeLock and InsideLock variable declarations produce the same compiler output:
// https://godbolt.org/z/eYWMPMTGz
// https://godbolt.org/z/7f6Ksdjeq

package reuse_test

import (
	"sync"
	"testing"
)

func BenchmarkDeclLock(b *testing.B) {

	b.Run("BeforeLock", func(b *testing.B) {
		m := sync.Mutex{}
		for i := 0; i < b.N; i++ {
			go func() {
				var key int //lint:ignore S1021
				m.Lock()
				key = i
				m.Unlock()
				_ = key
			}()
		}
	})

	b.Run("InsideLock", func(b *testing.B) {
		m := sync.Mutex{}
		for i := 0; i < b.N; i++ {
			go func() {
				m.Lock()
				key := i
				m.Unlock()
				_ = key
			}()
		}
	})

}
