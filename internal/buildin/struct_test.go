package buildin

import (
	"testing"
	"unsafe"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStructSize(t *testing.T) {
	Convey("基本数据类型大小检测", t, func() {
		So(unsafe.Sizeof(true), ShouldEqual, 1)
		So(unsafe.Sizeof(int8(0)), ShouldEqual, 1)
		So(unsafe.Sizeof(int16(0)), ShouldEqual, 2)
		So(unsafe.Sizeof(int32(0)), ShouldEqual, 4)
		So(unsafe.Sizeof(int64(0)), ShouldEqual, 8)
		So(unsafe.Sizeof(int(0)), ShouldEqual, 8)
		So(unsafe.Sizeof(float32(0)), ShouldEqual, 4)
		So(unsafe.Sizeof(float64(0)), ShouldEqual, 8)
		So(unsafe.Sizeof(""), ShouldEqual, 16)
		So(unsafe.Sizeof("hello world"), ShouldEqual, 16)
		So(unsafe.Sizeof([]int{}), ShouldEqual, 24)
		So(unsafe.Sizeof([]int{1, 2, 3}), ShouldEqual, 24)
		So(unsafe.Sizeof([3]int{1, 2, 3}), ShouldEqual, 24)
		So(unsafe.Sizeof(map[string]string{}), ShouldEqual, 8)
		So(unsafe.Sizeof(map[string]string{"1": "one", "2": "two"}), ShouldEqual, 8)
		So(unsafe.Sizeof(struct{}{}), ShouldEqual, 0)
	})

	Convey("自定义类型大小检测", t, func() {
		// |x---|
		So(unsafe.Sizeof(struct {
			i8 int8
		}{}), ShouldEqual, 1)

		// |x---|xxxx|xx--|
		So(unsafe.Sizeof(struct {
			i8  int8
			i32 int32
			i16 int16
		}{}), ShouldEqual, 12)

		// |x-xx|xxxx|
		So(unsafe.Sizeof(struct {
			i8  int8
			i16 int16
			i32 int32
		}{}), ShouldEqual, 8)

		// |x---|xxxx|xx--|----|xxxx|xxxx|
		So(unsafe.Sizeof(struct {
			i8  int8
			i32 int32
			i16 int16
			i64 int64
		}{}), ShouldEqual, 24)

		// |x-xx|xxxx|xxxx|xxxx|
		So(unsafe.Sizeof(struct {
			i8  int8
			i16 int16
			i32 int32
			i64 int64
		}{}), ShouldEqual, 16)

		type I8 int8
		type I16 int16
		type I32 int32

		So(unsafe.Sizeof(struct {
			i8  I8
			i16 I16
			i32 I32
		}{}), ShouldEqual, 8)
	})
}
