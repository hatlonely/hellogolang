package buildin

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

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

func TestCond(t *testing.T) {
	Convey("test cond", t, func() {
		var wg sync.WaitGroup
		deadline := time.Now().Add(1 * time.Second)
		q := NewBlockingQueue(4)
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int) {
				rd := rand.New(rand.NewSource(time.Now().UnixNano()))
				for time.Now().Before(deadline) {
					k := rd.Int() % 100
					q.put(k)
					fmt.Println(i, "put", k)
				}
				wg.Done()
			}(i)
		}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int) {
				for time.Now().Before(deadline) {
					fmt.Println(i, "take", q.take())
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
}

type BlockingQueue struct {
	mutex    *sync.Mutex
	notFull  *sync.Cond
	notEmpty *sync.Cond
	start    int
	end      int
	capacity int
	vector   []interface{}
}

func NewBlockingQueue(capacity int) *BlockingQueue {
	var mutex sync.Mutex
	return &BlockingQueue{
		mutex:    &mutex,
		notFull:  sync.NewCond(&mutex),
		notEmpty: sync.NewCond(&mutex),
		start:    0,
		end:      0,
		capacity: capacity,
		vector:   make([]interface{}, capacity+1),
	}
}

func (q *BlockingQueue) isEmpty() bool {
	return q.start == q.end
}

func (q *BlockingQueue) Size() int {
	return (q.end - q.start + q.capacity + 1) % (q.capacity + 1)
}

func (q *BlockingQueue) isFull() bool {
	return q.Size() == q.capacity
}

func (q *BlockingQueue) put(v interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	for q.isFull() {
		q.notFull.Wait()
	}
	q.vector[q.end] = v
	q.end++
	q.end %= q.capacity + 1
	q.notEmpty.Signal()
}

func (q *BlockingQueue) take() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	for q.isEmpty() {
		q.notEmpty.Wait()
	}
	res := q.vector[q.start]
	q.start++
	q.start %= q.capacity + 1
	q.notFull.Signal()
	return res
}
