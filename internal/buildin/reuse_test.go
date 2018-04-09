package buildin

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReuseSlice(t *testing.T) {
	Convey("slice 重用", t, func() {
		Convey("对一个 slice 执行 append 操作", func() {
			si1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			si2 := si1
			si2 = append(si2, 0)
			Convey("重新分配内存", func() {
				header1 := (*reflect.SliceHeader)(unsafe.Pointer(&si1))
				header2 := (*reflect.SliceHeader)(unsafe.Pointer(&si2))
				fmt.Println(header1.Data)
				fmt.Println(header2.Data)
				So(header1.Data, ShouldNotEqual, header2.Data)
			})
		})

		Convey("一个 slice 和它的切片", func() {
			si1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			si2 := si1[:7]
			Convey("不重新分配内存", func() {
				header1 := (*reflect.SliceHeader)(unsafe.Pointer(&si1))
				header2 := (*reflect.SliceHeader)(unsafe.Pointer(&si2))
				fmt.Println(header1.Data)
				fmt.Println(header2.Data)
				So(header1.Data, ShouldEqual, header2.Data)
			})

			Convey("往切片里面 append 一个值", func() {
				si2 = append(si2, 10)
				Convey("改变了原 slice 的值", func() {
					header1 := (*reflect.SliceHeader)(unsafe.Pointer(&si1))
					header2 := (*reflect.SliceHeader)(unsafe.Pointer(&si2))
					fmt.Println(header1.Data)
					fmt.Println(header2.Data)
					So(header1.Data, ShouldEqual, header2.Data)
					So(si1[7], ShouldEqual, 10)
				})
			})
		})

		Convey("make 每次都会返回新地址", func() {
			si := make([]int, 10, 10)
			header := (*reflect.SliceHeader)(unsafe.Pointer(&si))
			fmt.Println(header.Data)
			si = make([]int, 100, 100)
			header = (*reflect.SliceHeader)(unsafe.Pointer(&si))
			fmt.Println(header.Data)
			si = make([]int, 1000, 1000)
			header = (*reflect.SliceHeader)(unsafe.Pointer(&si))
			fmt.Println(header.Data)
			si = make([]int, 100, 100)
			header = (*reflect.SliceHeader)(unsafe.Pointer(&si))
			fmt.Println(header.Data)
			si = make([]int, 200, 200)
			header = (*reflect.SliceHeader)(unsafe.Pointer(&si))
			fmt.Println(header.Data)
			si = make([]int, 1000, 1000)
			header = (*reflect.SliceHeader)(unsafe.Pointer(&si))
			fmt.Println(header.Data)
		})
	})
}

func TestReuseString(t *testing.T) {
	Convey("string 重用", t, func() {
		Convey("字符串常量", func() {
			str1 := "hello world"
			str2 := "hello world"
			Convey("地址相同", func() {
				header1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
				header2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))
				fmt.Println(header1.Data)
				fmt.Println(header2.Data)
				So(header1.Data, ShouldEqual, header2.Data)
			})
		})

		Convey("相同字符串的不同子串", func() {
			str1 := "hello world"[:6]
			str2 := "hello world"[:5]
			Convey("地址相同", func() {
				header1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
				header2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))
				fmt.Println(header1.Data, str1)
				fmt.Println(header2.Data, str2)
				So(str1, ShouldNotEqual, str2)
				So(header1.Data, ShouldEqual, header2.Data)
			})
		})

		Convey("不同字符串的相同子串", func() {
			str1 := "hello world"[:5]
			str2 := "hello golang"[:5]
			Convey("地址不同", func() {
				header1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
				header2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))
				fmt.Println(header1.Data, str1)
				fmt.Println(header2.Data, str2)
				So(str1, ShouldEqual, str2)
				So(header1.Data, ShouldNotEqual, header2.Data)
			})
		})
	})
}

func TestMapReuse(t *testing.T) {
	Convey("map 重用", t, func() {
		m1 := map[string]int{"one": 1, "two": 2}
		m2 := m1
		m2["one"] = 10
		m2["three"] = 3
		Println(m1)
		Println(m2)
		So(m1["one"], ShouldEqual, 10)
		So(m2["one"], ShouldEqual, 10)
		So(m1, ShouldResemble, m2)
		So(m1, ShouldEqual, m2)
	})
}
