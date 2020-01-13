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
					case ch <- n:
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
