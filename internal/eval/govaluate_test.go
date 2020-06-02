package eval

import (
	"testing"

	"github.com/Knetic/govaluate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGoValuateMap(t *testing.T) {
	Convey("TestGoValuate", t, func() {
		Convey("case1", func() {
			expr, err := govaluate.NewEvaluableExpression("(key1 * key2 / 100) >= 90")
			So(err, ShouldBeNil)
			val, err := expr.Evaluate(map[string]interface{}{
				"key1": 100,
				"key2": 99,
			})
			So(err, ShouldBeNil)
			So(val, ShouldBeTrue)
		})

		Convey("case2", func() {
			expr, err := govaluate.NewEvaluableExpression("key1 * key2 / 100")
			So(err, ShouldBeNil)
			val, err := expr.Evaluate(map[string]interface{}{
				"key1": 100,
				"key2": 99,
			})
			So(err, ShouldBeNil)
			So(val, ShouldEqual, 99)
		})
	})
}
