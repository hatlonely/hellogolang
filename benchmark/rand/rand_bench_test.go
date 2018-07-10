package rand

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	mrand "math/rand"
	"sync"
	"testing"

	"github.com/spaolacci/murmur3"
)

func myrand() uint64 {
	var buf = make([]byte, 8)
	crand.Read(buf)
	return binary.LittleEndian.Uint64(buf)
}

func TestRand(t *testing.T) {
	fmt.Println(myrand())
	fmt.Println(mrand.Uint64())
	fmt.Println(mrand.Uint64())
	fmt.Println(murmur3.Sum64([]byte("0a9599d1-0920-4e97-b8fe-8407ccf0f387")))
	t.Error()
}

func BenchmarkRandConcurrent(b *testing.B) {
	b.Run("math rand", func(b *testing.B) {
		benchmarkRandConcurrent(b, mrand.Uint64)
	})
	b.Run("crypto rand", func(b *testing.B) {
		benchmarkRandConcurrent(b, myrand)
	})
	b.Run("murmur3 hash", func(b *testing.B) {
		benchmarkRandConcurrent(b, func() uint64 {
			return murmur3.Sum64([]byte("0a9599d1-0920-4e97-b8fe-8407ccf0f387"))
		})
	})
}

func benchmarkRandConcurrent(b *testing.B, f func() uint64) {
	var wg sync.WaitGroup
	for m := 0; m < 20; m++ {
		wg.Add(1)
		go func() {
			for i := 0; i < b.N; i++ {
				f()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
