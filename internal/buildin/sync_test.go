package buildin

import (
	"fmt"
	"sync"
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
