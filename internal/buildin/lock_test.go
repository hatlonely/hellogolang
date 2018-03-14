package buildin

import (
	"strconv"
	"sync"
	"testing"
)

func BenchmarkRWLock(b *testing.B) {
	cache := make(map[string]string)
	mutex := sync.RWMutex{}
	b.RunParallel(func(pb *testing.PB) {
		count := 0
		for pb.Next() {
			count++
			str := strconv.Itoa(count)

			mutex.Lock()
			cache["key"+str] = str
			mutex.Unlock()
		}
	})
	b.RunParallel(func(pb *testing.PB) {
		count := 0
		for pb.Next() {
			count++
			str := strconv.Itoa(count)

			mutex.RLock()
			_ = cache["key"+str]
			mutex.RUnlock()
		}
	})
}
