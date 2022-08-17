package errors

import (
	"fmt"
	"io"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"k8s.io/apimachinery/pkg/util/errors"
)

func TestK8sIoErrors(t *testing.T) {
	Convey("TestK8sIoErrors", t, func() {
		err5 := &MyError{"error5"}
		errs := errors.NewAggregate([]error{
			fmt.Errorf("error1"),
			errors.NewAggregate([]error{
				io.EOF,
				fmt.Errorf("error3"),
			}),
			fmt.Errorf("error4"),
			err5,
		})

		So(errs.Error(), ShouldEqual, "[error1, EOF, error3, error4, error5]")
		So(fmt.Sprint(errs.Errors()), ShouldEqual, "[error1 [EOF, error3] error4 error5]")
		So(fmt.Sprint(errors.Flatten(errs).Errors()), ShouldEqual, "[error1 EOF error3 error4 error5]")
		So(errs.Is(io.EOF), ShouldBeTrue)
		So(errs.Is(err5), ShouldBeTrue)

		So(errors.FilterOut(errs, func(err error) bool {
			if _, ok := err.(*MyError); ok {
				return true
			}
			return false
		}).Error(), ShouldEqual, "[error1, EOF, error3, error4]")
	})
}
