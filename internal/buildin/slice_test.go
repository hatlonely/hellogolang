package buildin

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSliceOP(t *testing.T) {
	Convey("Given 一个数组", t, func() {
		var arr []int

		Convey("When 初始化数组", func() {
			arr = make([]int, 0, 10)
			So(len(arr), ShouldEqual, 0)
			So(cap(arr), ShouldEqual, 10)

			arr = make([]int, 10)
			So(len(arr), ShouldEqual, 10)
			So(cap(arr), ShouldEqual, 10)
		})

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

		Convey("When 清空数组", func() {
			arr = []int{0, 1, 2, 3, 4}
			arr = arr[:0]
			Convey("Then 数组为空，容量不变", func() {
				So(len(arr), ShouldEqual, 0)
				So(cap(arr), ShouldEqual, 5)
			})
		})

		Convey("When 复制数组", func() {
			arr = []int{0, 1, 2, 3, 4}
			arrCopy := make([]int, 5)
			copy(arrCopy, arr)
			Convey("Then 数组应该相同", func() {
				So(arrCopy, ShouldResemble, []int{0, 1, 2, 3, 4})
			})
		})
	})
}

// 下面这个场景是关于slice的性能测试
// 关于这个问题的讨论见：http://hatlonely.github.io/2018/01/18/golang%20slice%E6%80%A7%E8%83%BD%E5%88%86%E6%9E%90/

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
		arr := make([]int, 0, N/10)
		for i := 0; i < N; i++ {
			arr = append(arr, i)
		}
	}
}

func BenchmarkAppendWithCapLessLen3th(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0, N/3)
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
		arr := make([]int, 0, N*10)
		for i := 0; i < N; i++ {
			arr = append(arr, i)
		}
	}
}

func BenchmarkAppendWithoutCapacityReuse(b *testing.B) {
	var arr []int
	for i := 0; i < b.N; i++ {
		arr = arr[:0]
		for i := 0; i < N; i++ {
			arr = append(arr, i)
		}
	}
}

func BenchmarkAppendWithCapEqualLenReuse(b *testing.B) {
	arr := make([]int, N)
	for i := 0; i < b.N; i++ {
		arr = arr[:0]
		for i := 0; i < N; i++ {
			arr = append(arr, i)
		}
	}
}

func BenchmarkAppendWithCapGreaterLen10thReuse(b *testing.B) {
	arr := make([]int, N*10)
	for i := 0; i < b.N; i++ {
		arr = arr[:0]
		for i := 0; i < N; i++ {
			arr = append(arr, i)
		}
	}
}
