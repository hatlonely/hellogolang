package retry_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	. "github.com/agiledragon/gomonkey"
	"github.com/giantswarm/retry-go"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRetryGo1(t *testing.T) {
	Convey("TestRetryGo", t, func() {
		patches := ApplyFuncSeq(DoSomething, []OutputCell{
			{Values: Params{0, errors.New("connection refuse")}, Times: 3},
			{Values: Params{123, nil}},
		})
		defer patches.Reset()

		var res int

		err := retry.Do(
			func() error {
				i, err := DoSomething()
				if err != nil {
					return err
				}

				res = i
				return nil
			},
			retry.Sleep(time.Second), // 重试间隔
			retry.MaxTries(4),        // 重试次数
			retry.RetryChecker(func(err error) bool {
				return err.Error() == "connection refuse"
			}), // 满足条件才重试
			retry.Timeout(5*time.Second), // 总共的超时时间
			retry.AfterRetry(func(err error) {
				fmt.Printf("[%v] err: [%v]\n", time.Now().Format(time.RFC3339), err)
			}), // 每次重试结束后执行
			retry.AfterRetryLimit(func(err error) {
				fmt.Printf("[%v] retry done, err: [%v]", time.Now().Format(time.RFC3339), err)
			}), // 全部重试完成后执行
		)

		So(err, ShouldBeNil)
		So(res, ShouldEqual, 123)
	})
}
