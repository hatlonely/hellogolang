package hystrix

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

var count int64

func myfunc() error {
	i := atomic.AddInt64(&count, 1)
	if i%2 == 0 {
		return fmt.Errorf("something wrong")
	}
	return nil
}

func TestHystrix(t *testing.T) {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:                200, // 超时时间 200ms
		MaxConcurrentRequests:  2,   // 最大并发数，超过并发返回错误
		RequestVolumeThreshold: 4,   // 请求数量的阀值，用这些数量的请求来计算阀值
		ErrorPercentThreshold:  25,  // 错误数量阀值，达到阀值，启动熔断
		SleepWindow:            300, // 熔断尝试恢复时间
	})

	var wg sync.WaitGroup
	for n := 0; n < 2; n++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				err := hystrix.Do("my_command", func() error {
					time.Sleep(time.Duration(100) * time.Millisecond)
					return myfunc()
				}, func(err error) error {
					time.Sleep(time.Duration(100) * time.Millisecond)
					return err
				})
				fmt.Println(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
