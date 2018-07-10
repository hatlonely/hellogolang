package bench

import (
	crand "crypto/rand"
	"encoding/base64"
	mrand "math/rand"
	"sync"
	"testing"

	"github.com/spaolacci/murmur3"
)

func randstr() []byte {
	buf := make([]byte, mrand.Intn(32)+8)
	crand.Read(buf)
	return buf
}

func base64encode(buf []byte) string {
	return base64.StdEncoding.EncodeToString(buf)
}

func murmur3hash(str string) uint64 {
	return murmur3.Sum64([]byte(str))
}

func fun() {
	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				murmur3hash(base64encode(randstr()))
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fun()
	}
}
