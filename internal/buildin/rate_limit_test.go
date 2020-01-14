package buildin

import (
	"context"
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/time/rate"
)

func TestRateLimit(t *testing.T) {
	Convey("test allow", t, func() {
		// 每 100ms 往桶里面 push 1个 token, token 的容量为 100，初始就是满的
		limiter := rate.NewLimiter(rate.Every(200*time.Millisecond), 100)
		So(limiter.Burst(), ShouldEqual, 100)
		So(limiter.Limit(), ShouldAlmostEqual, time.Second/(200*time.Millisecond))

		So(limiter.Allow(), ShouldBeTrue)                 // 取一个 token
		So(limiter.AllowN(time.Now(), 50), ShouldBeTrue)  // 取 50 个 token
		So(limiter.AllowN(time.Now(), 50), ShouldBeFalse) // 取 50 个 token，剩 49 个，不够，返回 false
	})

	Convey("test reserve", t, func() {
		limiter := rate.NewLimiter(rate.Every(200*time.Millisecond), 100)
		So(limiter.AllowN(time.Now(), 100), ShouldBeTrue)

		r := limiter.ReserveN(time.Now(), 10)
		So(r.OK(), ShouldBeTrue)
		time.Sleep(r.Delay())
		So(limiter.AllowN(time.Now(), 10), ShouldBeFalse)
	})

	Convey("test wait", t, func() {
		limiter := rate.NewLimiter(rate.Every(200*time.Millisecond), 100)
		now := time.Now()
		count := 0

		for {
			if time.Since(now) > 4800*time.Millisecond {
				break
			}
			// 阻塞等待一个从中取一个 token
			if err := limiter.Wait(context.Background()); err != nil {
				panic(err)
			}
			// 阻塞等待一个从中取 n 个 token
			if err := limiter.WaitN(context.Background(), 2); err != nil {
				panic(err)
			}
			count++
			fmt.Println("Got it", time.Now(), limiter.Burst(), limiter.Limit())
		}
		So(count, ShouldEqual, (100+4800/200)/3+1)
	})
}
