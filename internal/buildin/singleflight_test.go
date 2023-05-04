package buildin

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

func TestSingleFlightDo(t *testing.T) {
	var sf singleflight.Group

	for i := 0; i < 10; i++ {
		res, err, shared := sf.Do("key", func() (interface{}, error) {
			return fmt.Sprintf("value%d", i), nil
		})
		fmt.Println(res, err, shared)
	}
}

func TestSingleFlightDoParallel(t *testing.T) {
	var wg sync.WaitGroup
	var sf singleflight.Group

	ress := make([]interface{}, 10)
	errs := make([]interface{}, 10)
	shared := make([]bool, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			ress[i], errs[i], shared[i] = sf.Do("key", func() (interface{}, error) {
				time.Sleep(2 * time.Second)
				return fmt.Sprintf("value%d", i), nil
			})
			wg.Done()
		}(i)
	}

	wg.Wait()
	for i := 0; i < 10; i++ {
		fmt.Println(ress[i], errs[i], shared[i])
	}
}

func TestSingleFlightDoChan(t *testing.T) {
	var sf singleflight.Group
	chs := make([]<-chan singleflight.Result, 10)

	for i := 0; i < 10; i++ {
		chs[i] = sf.DoChan("key", func() (interface{}, error) {
			time.Sleep(2 * time.Second)
			return "value", nil
		})
	}

	for i := 0; i < 10; i++ {
		res := <-chs[i]
		fmt.Println(res.Val, res.Err, res.Shared)
	}
}

func TestSingleFlightForget(t *testing.T) {
	var sf singleflight.Group

	// 构造一个超时的请求
	ch1 := sf.DoChan("key", func() (interface{}, error) {
		time.Sleep(2 * time.Second)
		return nil, errors.New("timeout")
	})

	// 之后的请求发现已有一个请求，不会发起新的请求，从而都会超时
	for i := 0; i < 10; i++ {
		ch2 := sf.DoChan("key", func() (interface{}, error) {
			return "value", nil
		})
		select {
		case res := <-ch2:
			fmt.Println(res.Val, res.Err, res.Shared)
		case <-time.After(100 * time.Millisecond):
			fmt.Println("timeout")
		}
	}

	// 忘记前面的 key
	sf.Forget("key")

	// 之后的请求将重新发起新的请求，从而拿到正确的结果
	for i := 0; i < 10; i++ {
		ch2 := sf.DoChan("key", func() (interface{}, error) {
			return "value", nil
		})
		select {
		case res := <-ch2:
			fmt.Println(res.Val, res.Err, res.Shared)
		case <-time.After(100 * time.Millisecond):
			fmt.Println("timeout")
		}
	}

	// 最初的请求在函数体返回之后，返回超时
	res := <-ch1
	fmt.Println(res.Val, res.Err, res.Shared)
}
