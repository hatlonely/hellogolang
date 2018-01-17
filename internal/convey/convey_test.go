package convey

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec1(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}

func TestSpec2(t *testing.T) {
	Convey("Given 一个整数初始值", t, func() {
		x := 1

		Convey("When 这个整数自增", func() {
			x++

			Convey("Then 这个数比之前大1", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}
