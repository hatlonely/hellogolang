package syncmap

import (
	"math/rand"
	"sync"
	"testing"
)

func benchmarkMap(b *testing.B, hm Map) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for i := 0; i < 1; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100000; i++ {
					hm.Set(rand.Intn(1000), i*i)
					hm.Set(rand.Intn(1000), i*i)
					hm.Del(rand.Intn(1000))
				}
				wg.Done()
			}()
		}
		for i := 0; i < 80; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100000; i++ {
					hm.Get(i)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncmap(b *testing.B) {
	b.Run("mutex syncmap 1", func(b *testing.B) {
		hm := NewMutexSyncMap(1, func(key interface{}) int {
			return key.(int)
		})
		benchmarkMap(b, hm)
	})

	b.Run("mutex syncmap 100", func(b *testing.B) {
		hm := NewMutexSyncMap(100, func(key interface{}) int {
			return key.(int)
		})
		benchmarkMap(b, hm)
	})

	b.Run("std syncmap", func(b *testing.B) {
		hm := NewStdSyncMap()
		benchmarkMap(b, hm)
	})
}
