package unittest

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConvey1(t *testing.T) {
	Convey("TestConvey1", t, func() {
		So(1, ShouldEqual, 1)
	})
}

func TestConvey2(t *testing.T) {
	Convey("TestConvey2", t, func() {
		a := 1

		Convey("case1", func() {
			a += 3
			So(a, ShouldEqual, 4)
		})

		Convey("case2", func() {
			a += 4
			So(a, ShouldEqual, 5)
		})
	})
}
