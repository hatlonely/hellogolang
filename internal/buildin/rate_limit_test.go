package buildin

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestRateLimit(t *testing.T) {
	// 每 100ms 往桶里面 push 1个 token, token 的容量为 100，初始就是满的
	limiter := rate.NewLimiter(rate.Every(time.Duration(100)*time.Millisecond), 100)
	now := time.Now()
	count := 0
	for {
		if time.Since(now) > time.Duration(4800)*time.Millisecond {
			break
		}
		// 阻塞等待一个从中取一个 token
		if err := limiter.Wait(context.Background()); err != nil {
			panic(err)
		}
		count++
		fmt.Println("Got it", time.Now())
	}
	fmt.Println(count)
}
