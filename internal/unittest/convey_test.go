package unittest

import (
	"fmt"
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

func TestConvey3(t *testing.T) {
	Convey("TestConvey3", t, func() {
		fmt.Println("enter layer 1")

		fmt.Println("start layer 2-1")
		Convey("layer 2-1", func() {
			fmt.Println("enter layer 2-1")

			fmt.Println("start layer 3-1")
			Convey("layer 3-1", func() {
				fmt.Println("enter layer 3-1")
			})
			fmt.Println("end layer 3-1")

			fmt.Println("start layer 3-2")
			Convey("layer 3-2", func() {
				fmt.Println("enter layer 3-2")
			})
			fmt.Println("end layer 3-2")
		})
		fmt.Println("end layer 2-1")

		fmt.Println("start layer 2-2")
		Convey("layer 2-2", func() {
			fmt.Println("enter layer 2-2")
		})
		fmt.Println("end layer 2-2")
	})
}
