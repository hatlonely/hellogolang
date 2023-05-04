package buildin

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

func TestSingleFight1(t *testing.T) {
	var sf singleflight.Group

	for i := 0; i < 10; i++ {
		res, err, shared := sf.Do("testSingleFlight", func() (interface{}, error) {
			t.Log("hello world")
			return "testSingleFlight", nil
		})
		fmt.Println(res, err, shared)
	}
}

func TestSingleFightParallel(t *testing.T) {
	var wg sync.WaitGroup
	var sf singleflight.Group

	ress := make([]interface{}, 10, 10)
	errs := make([]interface{}, 10, 10)
	shared := make([]bool, 10, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			ress[i], errs[i], shared[i] = sf.Do("testSingleFlight", func() (interface{}, error) {
				t.Log("hello world")
				time.Sleep(2)
				return "testSingleFlight", nil
			})
			wg.Done()
		}(i)
	}

	wg.Wait()
	for i := 0; i < 10; i++ {
		fmt.Println(ress[i], errs[i], shared[i])
	}
}
