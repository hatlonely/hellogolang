package syncmap

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/cornelk/hashmap"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHashMap(t *testing.T) {
	Convey("test hashmap", t, func() {
		hm := hashmap.HashMap{}
		hm.Set(1, 2)
		val, ok := hm.Get(1)
		So(ok, ShouldBeTrue)
		So(val, ShouldEqual, 2)

		for j := 0; j < 1000; j++ {
			// hm.Set(rand.Intn(100), i*100+j)
			// this lib is not supprt del operation
			hm.Set(j, j)
		}
		hm.Del(rand.Intn(100))

		// var wg sync.WaitGroup
		// for i := 0; i < 1; i++ {
		// 	wg.Add(1)
		// 	go func(i int) {
		// 		for j := 0; j < 100; j++ {
		// 			// hm.Set(rand.Intn(100), i*100+j)
		// 			// this lib is not supprt del operation
		// 			hm.Del(rand.Intn(100))
		// 		}
		// 		wg.Done()
		// 	}(i)
		// }
		// wg.Wait()
	})
}

func TestLockfreeMap(t *testing.T) {
	Convey("test lockfree map", t, func() {
		hm := NewLockfreeMap(30, func(i interface{}) int {
			return i.(int)
		})

		Convey("test set/get", func() {
			for i := 0; i < 100; i++ {
				val, ok := hm.Get(i)
				So(ok, ShouldBeFalse)
				So(val, ShouldEqual, nil)
			}
			for i := 0; i < 100; i++ {
				hm.Set(i, i*i)
			}
			time.Sleep(time.Millisecond)
			for i := 0; i < 100; i++ {
				val, ok := hm.Get(i)
				So(val, ShouldEqual, i*i)
				So(ok, ShouldBeTrue)
			}
			// delete 33~60
			for i := 33; i < 60; i++ {
				hm.Del(i)
			}
			time.Sleep(time.Millisecond)
			for i := 0; i < 33; i++ {
				val, ok := hm.Get(i)
				So(val, ShouldEqual, i*i)
				So(ok, ShouldBeTrue)
			}
			for i := 33; i < 60; i++ {
				val, ok := hm.Get(i)
				So(ok, ShouldBeFalse)
				So(val, ShouldEqual, nil)
			}
			for i := 60; i < 100; i++ {
				val, ok := hm.Get(i)
				So(val, ShouldEqual, i*i)
				So(ok, ShouldBeTrue)
			}
			// delete 0~10
			for i := 0; i < 10; i++ {
				hm.Del(i)
			}
			time.Sleep(time.Millisecond)
			for i := 0; i < 10; i++ {
				val, ok := hm.Get(i)
				So(ok, ShouldBeFalse)
				So(val, ShouldEqual, nil)
			}
			for i := 10; i < 33; i++ {
				val, ok := hm.Get(i)
				So(val, ShouldEqual, i*i)
				So(ok, ShouldBeTrue)
			}
			for i := 33; i < 60; i++ {
				val, ok := hm.Get(i)
				So(ok, ShouldBeFalse)
				So(val, ShouldEqual, nil)
			}
			for i := 60; i < 100; i++ {
				val, ok := hm.Get(i)
				So(val, ShouldEqual, i*i)
				So(ok, ShouldBeTrue)
			}
			// delete 87~100
			for i := 0; i < 10; i++ {
				val, ok := hm.Get(i)
				So(ok, ShouldBeFalse)
				So(val, ShouldEqual, nil)
			}
			for i := 87; i < 100; i++ {
				hm.Del(i)
			}
			time.Sleep(time.Millisecond)
			for i := 10; i < 33; i++ {
				val, ok := hm.Get(i)
				So(val, ShouldEqual, i*i)
				So(ok, ShouldBeTrue)
			}
			for i := 33; i < 60; i++ {
				val, ok := hm.Get(i)
				So(ok, ShouldBeFalse)
				So(val, ShouldEqual, nil)
			}
			for i := 60; i < 87; i++ {
				val, ok := hm.Get(i)
				So(val, ShouldEqual, i*i)
				So(ok, ShouldBeTrue)
			}
			for i := 87; i < 100; i++ {
				val, ok := hm.Get(i)
				So(ok, ShouldBeFalse)
				So(val, ShouldEqual, nil)
			}
			// hm.show()
			// set 1~100
			for i := 0; i < 100; i++ {
				hm.Set(i, i+1)
			}
			time.Sleep(time.Millisecond)
			// hm.show()
			for i := 0; i < 100; i++ {
				val, ok := hm.Get(i)
				So(ok, ShouldBeTrue)
				So(val, ShouldEqual, i+1)
			}
		})

		Convey("test concurrent", func() {
			var wg sync.WaitGroup
			wg.Add(1)
			// hm := hashmap.HashMap{}
			go func() {
				for j := 0; j < 100000; j++ {
					hm.Del(rand.Intn(100))
					hm.Set(rand.Intn(100), rand.Intn(1000))
					hm.Set(rand.Intn(100), rand.Intn(1000))
					hm.Set(rand.Intn(100), rand.Intn(1000))
				}
				wg.Done()
			}()
			for i := 0; i < 100; i++ {
				wg.Add(1)
				go func(i int) {
					for j := 0; j < 100000; j++ {
						hm.Del(rand.Intn(100))
						hm.Set(rand.Intn(100), rand.Intn(1000))
						hm.Set(rand.Intn(100), rand.Intn(1000))
						hm.Set(rand.Intn(100), rand.Intn(1000))
						hm.Get(i)
					}
					wg.Done()
				}(i)
			}
			wg.Wait()
			// hm.show()
			t.Error()
		})
	})
}

func BenchmarkHashMap(b *testing.B) {
	b.Run("hashmap set ", func(b *testing.B) {
		hm := hashmap.HashMap{}
		for i := 0; i < b.N; i++ {
			hm.Set(1, 2)
		}
	})

	b.Run("stdmap set", func(b *testing.B) {
		hm := map[int]int{}
		for i := 0; i < b.N; i++ {
			hm[1] = 2
		}
	})

	b.Run("syncmap set", func(b *testing.B) {
		hm := sync.Map{}
		for i := 0; i < b.N; i++ {
			hm.Store(1, 2)
		}
	})

	b.Run("lockfree set", func(b *testing.B) {
		hm := NewLockfreeMap(100, func(key interface{}) int {
			return key.(int)
		})
		for i := 0; i < b.N; i++ {
			hm.Set(1, 2)
		}
	})

	b.Run("my sync map get ", func(b *testing.B) {
		hm := NewSyncMap(20, func(key interface{}) int {
			return key.(int)
		})
		hm.Set(1, 2)
		for i := 0; i < b.N; i++ {
			_, _ = hm.Get(1)
		}
	})

	b.Run("hashmap get ", func(b *testing.B) {
		hm := hashmap.HashMap{}
		hm.Set(1, 2)
		for i := 0; i < b.N; i++ {
			_, _ = hm.Get(1)
		}
	})

	b.Run("stdmap get", func(b *testing.B) {
		hm := map[int]int{}
		hm[1] = 2
		for i := 0; i < b.N; i++ {
			_, _ = hm[1]
		}
	})

	b.Run("syncmap get", func(b *testing.B) {
		hm := sync.Map{}
		hm.Store(1, 2)
		for i := 0; i < b.N; i++ {
			hm.Load(1)
		}
	})

	b.Run("my sync map get", func(b *testing.B) {
		hm := NewSyncMap(20, func(key interface{}) int {
			return key.(int)
		})
		hm.Set(1, 2)
		for i := 0; i < b.N; i++ {
			hm.Get(1)
		}
	})

	b.Run("lockfree get", func(b *testing.B) {
		hm := NewLockfreeMap(100, func(key interface{}) int {
			return key.(int)
		})
		hm.Set(1, 2)
		for i := 0; i < b.N; i++ {
			hm.Get(1)
		}
	})
}

type Map interface {
	Set(key interface{}, val interface{})
	Get(key interface{}) (interface{}, bool)
	Del(key interface{})
}

func benchmarkMap(b *testing.B, hm Map) {
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
	// for i := 0; i < 1; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		for i := 0; i < 100000; i++ {
	// 			hm.Del(rand.Intn(1000))
	// 		}
	// 		wg.Done()
	// 	}()
	// }
	for i := 0; i < 100; i++ {
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

func BenchmarkMultiThread(b *testing.B) {
	b.Run("lockfree set/get/del", func(b *testing.B) {
		hm := NewLockfreeMap(100, func(key interface{}) int {
			return key.(int)
		})
		benchmarkMap(b, hm)
	})

	b.Run("my syncmap set/get/del-100", func(b *testing.B) {
		hm := NewSyncMap(100, func(key interface{}) int {
			return key.(int)
		})
		benchmarkMap(b, hm)
	})

	b.Run("my syncmap set/get/del-1", func(b *testing.B) {
		hm := NewSyncMap(1, func(key interface{}) int {
			return key.(int)
		})
		benchmarkMap(b, hm)
	})

	b.Run("std syncmap set/get/del", func(b *testing.B) {
		hm := &StdSyncMap{}
		benchmarkMap(b, hm)
	})
}
