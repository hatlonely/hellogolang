package retry_test

import (
	"errors"
	"testing"

	. "github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/matryer/try.v1"
)

func TestTry(t *testing.T) {
	Convey("TestRetryGo", t, func() {
		patches := ApplyFuncSeq(DoSomething, []OutputCell{
			{Values: Params{0, errors.New("connection refuse")}, Times: 2},
			{Values: Params{123, nil}},
		})
		defer patches.Reset()

		var res int
		err := try.Do(func(attempt int) (bool, error) {
			i, err := DoSomething()
			if err != nil {
				return attempt < 3, err
			}
			res = i

			return false, nil
		})

		So(err, ShouldBeNil)
		So(res, ShouldEqual, 123)
	})
}
