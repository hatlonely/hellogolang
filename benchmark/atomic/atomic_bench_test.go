package atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

var i = int64(0)
var mutex sync.Mutex

func mutexInc() {
	mutex.Lock()
	defer mutex.Unlock()
	i++
}

var rwMutex sync.RWMutex

func rwMutexInc() {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	i++
}

func atomicInc() {
	atomic.AddInt64(&i, 1)
}

func inc() {
	i++
}

func benchmark(b *testing.B, f func()) {
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

func BenchmarkAtomic(b *testing.B) {
	b.Run("void", func(b *testing.B) {
		benchmark(b, func() {})
	})
	b.Run("inc", func(b *testing.B) {
		benchmark(b, func() { i++ })
	})
	b.Run("mutex", func(b *testing.B) {
		benchmark(b, mutexInc)
	})
	b.Run("rwmutex", func(b *testing.B) {
		benchmark(b, rwMutexInc)
	})
	b.Run("atomic", func(b *testing.B) {
		benchmark(b, atomicInc)
	})
}
