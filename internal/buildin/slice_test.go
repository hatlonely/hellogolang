package buildin

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSliceOP(t *testing.T) {
	Convey("Given 一个数组", t, func() {
		var arr []int

		Convey("When 插入元素", func() {
			for i := 0; i < 5; i++ {
				arr = append(arr, i)
			}

			Convey("Then 数组里面必须包含这些元素", func() {
				So(len(arr), ShouldEqual, 5)
				So(arr, ShouldResemble, []int{0, 1, 2, 3, 4})
			})
		})

		Convey("When 使用切片", func() {
			arr = []int{0, 1, 2, 3, 4}
			sli := arr[2:4]

			Convey("Then 切片的结果正确", func() {
				So(len(sli), ShouldEqual, 2)
				So(sli, ShouldResemble, []int{2, 3})
			})
		})

		Convey("When 删除元素", func() {
			arr = []int{0, 1, 2, 3, 4}
		})
	})
}

var N = 3000000

func BenchmarkAppendWithoutCapacity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var arr []int
		for i := 0; i < N; i++ {
			arr = append(arr, i)
		}
	}
}

func BenchmarkAppendWithCapLessLen10th(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0, N / 10)
		for i := 0; i < N; i++ {
			arr = append(arr, i)
		}
	}
}

func BenchmarkAppendWithCapLessLen3th(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0, N / 3)
		for i := 0; i < N; i++ {
			arr = append(arr, i)
		}
	}
}

func BenchmarkAppendWithCapEqualLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0, N)
		for i := 0; i < N; i++ {
			arr = append(arr, i)
		}
	}
}

func BenchmarkAppendWithCapGreaterLen10th(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0, N * 10)
		for i := 0; i < N; i++ {
			arr = append(arr, i)
		}
	}
}
