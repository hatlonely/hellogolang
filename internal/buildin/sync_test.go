package buildin

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMutex(t *testing.T) {
	Convey("test mutex", t, func() {
		var mutex sync.Mutex
		for i := 0; i < 5; i++ {
			go func() {
				mutex.Lock()
				fmt.Println("hello world")
				mutex.Unlock()
			}()
		}
	})
}

func TestRWMutex(t *testing.T) {
	Convey("test rwmutex", t, func() {
		var mutex sync.RWMutex
		kvs := map[int]int{}
		for i := 0; i < 10; i++ {
			go func() {
				for j := 0; j < 100; j++ {
					mutex.Lock()
					kvs[j] = i
					mutex.Unlock()
				}
			}()
		}

		for i := 0; i < 10; i++ {
			go func() {
				for j := 0; j < 100; j++ {
					mutex.RLock()
					_ = kvs[j]
					mutex.RUnlock()
				}
			}()
		}
	})
}

func TestSyncMap(t *testing.T) {
	Convey("test sync map", t, func() {
		kvs := &sync.Map{}

		kvs.Store("key1", "val1")
		val1, ok1 := kvs.Load("key1")
		So(ok1, ShouldBeTrue)
		So(val1, ShouldEqual, "val1")

		val2, ok2 := kvs.Load("key2")
		So(ok2, ShouldBeFalse)
		So(val2, ShouldBeNil)

		kvs.Delete("key1")

		val3, ok3 := kvs.LoadOrStore("key3", "val3")
		So(ok3, ShouldBeFalse)
		So(val3, ShouldEqual, "val3")

		val3, ok3 = kvs.LoadOrStore("key3", "val33")
		So(ok3, ShouldBeTrue)
		So(val3, ShouldEqual, "val3")

		kvs.Range(func(key, val interface{}) bool {
			fmt.Println(key, val)
			return true
		})
	})

	Convey("test concurrent sync map", t, func() {
		kvs := &sync.Map{}
		for i := 0; i < 10; i++ {
			go func() {
				for j := 0; j < 100; j++ {
					kvs.Store(j, i)
				}
			}()
		}

		for i := 0; i < 10; i++ {
			go func() {
				for j := 0; j < 100; j++ {
					_, _ = kvs.Load(j)
				}
			}()
		}

		for i := 0; i < 10; i++ {
			go func() {
				for j := 0; j < 100; j++ {
					kvs.Delete(j)
				}
			}()
		}
	})
}

func TestWaitGroup(t *testing.T) {
	Convey("test wait group", t, func() {
		var wg sync.WaitGroup

		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func() {
				for j := 0; j < 100; j++ {
					_ = strconv.Itoa(j)
				}
				wg.Done()
			}()
		}

		wg.Wait()
	})
}

func TestPool(t *testing.T) {
	Convey("test pool", t, func() {
		pool := &sync.Pool{
			New: func() interface{} {
				return 0
			},
		}

		pool.Put(1)

		So(pool.Get(), ShouldEqual, 1)
	})
}

func TestOnce(t *testing.T) {
	Convey("test once", t, func() {
		{
			var wg sync.WaitGroup
			count := int64(0)
			for i := 0; i < 10; i++ {
				wg.Add(1)
				go func() {
					atomic.AddInt64(&count, 1)
					wg.Done()
				}()
			}
			wg.Wait()
			So(count, ShouldEqual, 10)
		}
		{
			var wg sync.WaitGroup
			once := sync.Once{}
			count := int64(0)
			for i := 0; i < 10; i++ {
				wg.Add(1)
				go func() {
					once.Do(func() {
						atomic.AddInt64(&count, 1)
						wg.Done()
					})
				}()
			}
			So(count, ShouldEqual, 1)
		}
	})
}
