package errors_test

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestErrors(t *testing.T) {
	err1 := errors.New("err1 message")
	err2 := errors.Wrap(err1, "err2 message")
	err3 := errors.Wrap(err2, "err3 message")

	fmt.Println(err3.Error())
	fmt.Printf("%v\n", errors.Cause(err3))
	fmt.Printf("%T\n", errors.Cause(err3))
	fmt.Printf("%+v\n", err3)
}

func TestConstructor(t *testing.T) {
	Convey("errors.New", t, func() {
		err := errors.New("timeout")
		So(err.Error(), ShouldEqual, "timeout")
		So(fmt.Sprint(err), ShouldEqual, "timeout")
		fmt.Printf("%+v", err) // 输出堆栈信息
	})

	Convey("errors.Errorf", t, func() {
		err := errors.Errorf("timeout")
		So(err.Error(), ShouldEqual, "timeout")
	})
}

func TestWithMessage(t *testing.T) {
	Convey("with message", t, func() {
		err0 := fmt.Errorf("timeout")
		err1 := errors.WithMessage(err0, "wrap1")
		err2 := errors.WithMessage(err1, "wrap2")

		So(err0.Error(), ShouldEqual, "timeout")
		So(err1.Error(), ShouldEqual, "wrap1: timeout")
		So(err2.Error(), ShouldEqual, "wrap2: wrap1: timeout")

		So(fmt.Sprintf("%+v", err0), ShouldEqual, "timeout")
		So(fmt.Sprintf("%+v", err1), ShouldEqual, "timeout\nwrap1")
		So(fmt.Sprintf("%+v", err2), ShouldEqual, "timeout\nwrap1\nwrap2")
	})
}

func TestWithStack(t *testing.T) {
	Convey("with stack", t, func() {
		err0 := fmt.Errorf("timeout")
		err1 := errors.WithStack(err0)

		So(err0.Error(), ShouldEqual, "timeout")
		So(err1.Error(), ShouldEqual, "timeout")

		So(fmt.Sprintf("%+v", err0), ShouldEqual, "timeout")
		So(fmt.Sprintf("%+v", err1), ShouldContainSubstring, `timeout
github.com/hatlonely/hellogolang/internal/errors_test.TestWithStack.func1
	/Users/hatlonely/hatlonely/github.com/hatlonely/hellogolang/internal/errors/errors_test.go:55`)
		fmt.Printf("%+v", err1) // 输出堆栈信息
	})
}

func TestWrap(t *testing.T) {
	Convey("wrap", t, func() {
		err0 := fmt.Errorf("timeout")
		err1 := errors.Wrap(err0, "wrap1")
		err2 := errors.Wrap(err1, "wrap2")

		So(err0.Error(), ShouldEqual, "timeout")
		So(err1.Error(), ShouldEqual, "wrap1: timeout")
		So(err2.Error(), ShouldEqual, "wrap2: wrap1: timeout")

		fmt.Printf("%+v\n", err1)
		fmt.Printf("%+v\n", err2)

	})
}
