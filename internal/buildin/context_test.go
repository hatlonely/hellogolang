package buildin

import (
	"context"
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContext(t *testing.T) {
	Convey("test context cancel", t, func() {
		generator := func(ctx context.Context) <-chan int {
			ch := make(chan int)
			n := 1
			go func() {
				for {
					select {
					case <-ctx.Done():
						return
					default:
						ch <- n
						n++
					}
				}
			}()
			return ch
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		for n := range generator(ctx) {
			fmt.Println(n)
			if n == 5 {
				break
			}
		}
	})

	Convey("test context deadline", t, func() {
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(50*time.Millisecond))
		defer cancel()

		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}
	})

	Convey("test context timeout", t, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()

		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}
	})

	Convey("test context value", t, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		ctx1 := context.WithValue(ctx, "key1", "val1")
		ctx2 := context.WithValue(ctx1, "key2", "val2")

		defer cancel()

		So(ctx2.Value("key1"), ShouldEqual, "val1")
		So(ctx2.Value("key2"), ShouldEqual, "val2")
		So(ctx2.Value("key3"), ShouldBeNil)
	})
}

func TestSelect(t *testing.T) {
	Convey("test select", t, func() {
		ch1 := make(chan int, 1)
		ch2 := make(chan int, 1)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

	out:
		for {
			// 多个 case 条件同时满足，执行顺序是不确定的
			// 没有一个 case 满足时，才会执行 default
			select {
			case i1 := <-ch1:
				fmt.Println(i1)
				time.Sleep(time.Second)
				ch1 <- 1
			case i2 := <-ch2:
				fmt.Println(i2)
				ch2 <- 2
				time.Sleep(time.Second)
			case <-ctx.Done():
				fmt.Println("done")
				time.Sleep(time.Second)
				break out
			default:
				fmt.Println("default")
				ch1 <- 1
				ch2 <- 2
				time.Sleep(time.Second)
			}
		}
	})
}
