package eval

import (
	"context"
	"strconv"
	"testing"

	"github.com/PaesslerAG/gval"
	"github.com/PaesslerAG/jsonpath"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGvalEvaluable(t *testing.T) {
	Convey("TestGvalEvaluable", t, func() {
		Convey("evaluable", func() {
			eval, err := gval.Full().NewEvaluable(`key1 * (key2 + 1)`)
			So(err, ShouldBeNil)
			val, err := eval(context.Background(), map[string]interface{}{
				"key1": 2,
				"key2": 4,
			})
			So(err, ShouldBeNil)
			So(val, ShouldEqual, 10)
		})
		Convey("evaluable with function", func() {
			eval, err := gval.Full(
				gval.Constant("N", 10),
				gval.Function("strlen", func(str string) (int, error) {
					return len(str), nil
				}),
				gval.Function("toint", func(str string) (int, error) {
					return strconv.Atoi(str)
				}),
				jsonpath.Language(),
			).NewEvaluable(`strlen(key1) + N + toint(key2) + $["key-3"]`)

			So(err, ShouldBeNil)
			val, err := eval(context.Background(), map[string]interface{}{
				"key1":  "1234",
				"key2":  "12",
				"key-3": 4,
			})
			So(err, ShouldBeNil)
			So(val, ShouldEqual, 30)
		})
	})
}

func TestGvalEvaluate(t *testing.T) {
	Convey("TestGvalEvaluate", t, func() {
		Convey("map[string]interface", func() {
			val, err := gval.Evaluate(`key1 * (key2 + 1)`, map[string]interface{}{
				"key1": 2,
				"key2": 4,
			})
			So(err, ShouldBeNil)
			So(val, ShouldEqual, 10)
		})
		Convey("map[string]int", func() {
			val, err := gval.Evaluate(`key1 * (key2 + 1)`, map[string]int{
				"key1": 2,
				"key2": 4,
			})
			So(err, ShouldBeNil)
			So(val, ShouldEqual, 10)
		})
		Convey("map[string]string", func() {
			val, err := gval.Evaluate(`key1 * (key2 + 1)`, map[string]string{
				"key1": "2",
				"key2": "4",
			})
			So(err, ShouldBeNil)
			So(val, ShouldEqual, 10)
		})
		Convey("string + string => string", func() {
			val, err := gval.Evaluate(`key1 + key2`, map[string]string{
				"key1": "2",
				"key2": "4",
			})
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "24")
		})
		Convey("int + string-convert-int => int", func() {
			val, err := gval.Evaluate(`key1 + key2`, map[string]interface{}{
				"key1": 2,
				"key2": "4",
			})
			So(err, ShouldBeNil)
			So(val, ShouldEqual, 6)
		})
		Convey("int + string => string", func() {
			val, err := gval.Evaluate(`key1 + key2`, map[string]interface{}{
				"key1": 2,
				"key2": "4xx",
			})
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "24xx")
		})
		Convey("struct", func() {
			val, err := gval.Evaluate(`Key1 * (Key2.Key3 + 1)`, struct {
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
		Convey("json key", func() {
			val, err := gval.Evaluate(`$["key-a"] + 10`, map[string]interface{}{
				"key-a": 10,
			}, jsonpath.Language())
			So(err, ShouldBeNil)
			So(val, ShouldEqual, 20)
		})
		Convey("function", func() {
			val, err := gval.Evaluate(`strlen(key1) + 10 + toint(key2)`, map[string]interface{}{
				"key1": "1234",
				"key2": "12",
			}, gval.Function("strlen", func(str string) (int, error) {
				return len(str), nil
			}), gval.Function("toint", func(str string) (int, error) {
				return strconv.Atoi(str)
			}))
			So(err, ShouldBeNil)
			So(val, ShouldEqual, 26)
		})
	})
}
