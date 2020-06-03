package eval

import (
	"testing"

	"github.com/antonmedv/expr"
	. "github.com/smartystreets/goconvey/convey"
)

func TestExpr(t *testing.T) {
	Convey("TestExpr", t, func() {
		Convey("map[string]interface{}", func() {
			prg, err := expr.Compile(`key1 * (key2 + 1)`)
			So(err, ShouldBeNil)
			val, err := expr.Run(prg, map[string]interface{}{
				"key1": 2,
				"key2": 4,
			})
			So(err, ShouldBeNil)
			So(val, ShouldEqual, 10)
		})

		Convey("struct", func() {
			prg, err := expr.Compile(`Key1 * (Key2.Key3 + 1)`)
			So(err, ShouldBeNil)
			val, err := expr.Run(prg, struct {
				Key1 int
				Key2 struct {
					Key3 int
				}
			}{
				Key1: 2,
				Key2: struct{ Key3 int }{Key3: 4},
			})
			So(err, ShouldEqual, nil)
			So(val, ShouldEqual, 10)
		})

		Convey("eval", func() {
			val, err := expr.Eval(`key1 * (key2 + 1)`, map[string]interface{}{
				"key1": 2,
				"key2": 4,
			})
			So(err, ShouldBeNil)
			So(val, ShouldEqual, 10)
		})

		Convey("string cannot convert to int automatic", func() {
			_, err := expr.Eval(`key1 * (key2 + 1)`, map[string]string{
				"key1": "2",
				"key2": "4",
			})
			So(err, ShouldNotBeNil)
		})
	})
}
