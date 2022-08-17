package buildin

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestError_Fmt(t *testing.T) {
	Convey("fmt err", t, func() {
		err := fmt.Errorf("fmt err")
		So(err.Error(), ShouldEqual, "fmt err")
		So(fmt.Sprint(err), ShouldEqual, "fmt err")
		So(fmt.Sprintf("%+v", err), ShouldEqual, "fmt err")
	})
}

func TestError_Errors(t *testing.T) {
	Convey("TestError", t, func() {
		err := errors.New("errors err")
		So(err.Error(), ShouldEqual, "errors err")
		So(fmt.Sprint(err), ShouldEqual, "errors err")
		So(fmt.Sprintf("%+v", err), ShouldEqual, "errors err")
	})
}

func TestErrors_Wrap(t *testing.T) {
	Convey("TestErrors_Wrap", t, func() {
		err0 := errors.New("inner err")
		err1 := fmt.Errorf("wrap1: %w", err0)
		err2 := fmt.Errorf("wrap2: %w", err1)

		So(err0.Error(), ShouldEqual, "inner err")
		So(err1.Error(), ShouldEqual, "wrap1: inner err")
		So(err2.Error(), ShouldEqual, "wrap2: wrap1: inner err")

		So(fmt.Sprintf("%+v", err0), ShouldEqual, "inner err")
		So(fmt.Sprintf("%+v", err1), ShouldEqual, "wrap1: inner err")
		So(fmt.Sprintf("%+v", err2), ShouldEqual, "wrap2: wrap1: inner err")

		So(errors.Unwrap(err0), ShouldBeNil)
		So(errors.Unwrap(err1), ShouldEqual, err0)
		So(errors.Unwrap(err2), ShouldEqual, err1)
	})
}

type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return e.message
}

func TestErrors_IsAs(t *testing.T) {
	Convey("TestErrors_Wrap", t, func() {
		err0 := &MyError{message: "inner err"}
		err1 := fmt.Errorf("wrap1: %w", err0)
		err2 := fmt.Errorf("wrap2: %w", err1)

		So(err0.Error(), ShouldEqual, "inner err")
		So(err1.Error(), ShouldEqual, "wrap1: inner err")
		So(err2.Error(), ShouldEqual, "wrap2: wrap1: inner err")

		So(errors.Unwrap(err0), ShouldBeNil)
		So(errors.Unwrap(err1), ShouldEqual, err0)
		So(errors.Unwrap(err2), ShouldEqual, err1)

		So(errors.Is(err0, err0), ShouldBeTrue)
		So(errors.Is(err1, err0), ShouldBeTrue)
		So(errors.Is(err2, err0), ShouldBeTrue)

		var err *MyError
		So(errors.As(err0, &err), ShouldBeTrue)
		So(err, ShouldEqual, err0)
		So(errors.As(err1, &err), ShouldBeTrue)
		So(err, ShouldEqual, err0)
		So(errors.As(err2, &err), ShouldBeTrue)
		So(err, ShouldEqual, err0)
	})
}
