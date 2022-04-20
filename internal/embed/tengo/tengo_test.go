package tengo

import (
	"context"
	"fmt"
	"testing"

	"github.com/d5/tengo/v2"
	"github.com/smartystreets/goconvey/convey"
)

func TestTengoRun(t *testing.T) {
	convey.Convey("TestTengoRun", t, func() {
		script := tengo.NewScript([]byte(`
each := func(seq, fn) {
	for x in seq {
		fn(x)
	}
}

sum := 0
mul := 1

each([a, b, c, d], func(x) {
	sum += x
	mul *= x
})

`))
		_ = script.Add("a", 1)
		_ = script.Add("b", 9)
		_ = script.Add("c", 8)
		_ = script.Add("d", 4)

		compiled, err := script.RunContext(context.Background())
		convey.So(err, convey.ShouldBeNil)

		sum := compiled.Get("sum")
		mul := compiled.Get("mul")

		convey.So(sum.Int(), convey.ShouldEqual, 22)
		convey.So(mul.Int(), convey.ShouldEqual, 288)
	})
}

func TestTengoEval(t *testing.T) {
	convey.Convey("TestTengoEval", t, func() {
		res, err := tengo.Eval(context.Background(), `
	input ? "success" : "fail"
`, map[string]interface{}{
			"input": 1,
		})

		convey.So(err, convey.ShouldBeNil)

		convey.So(res, convey.ShouldEqual, "success")
	})
}

func TestTengoCallFunction(t *testing.T) {
	convey.Convey("TestTengoCall", t, func() {
		script := tengo.NewScript([]byte(`
mul := func(a, b) {
	return a * b
}
`))
		c, err := script.Compile()
		convey.So(err, convey.ShouldBeNil)
		convey.So(c.Run(), convey.ShouldBeNil)
		fmt.Println(c.GetAll())
	})
}
