package retry_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	. "github.com/agiledragon/gomonkey"
	"github.com/avast/retry-go"
	. "github.com/smartystreets/goconvey/convey"
)

func DoSomething() (int, error) {
	return 0, nil
}

func TestRetryGo(t *testing.T) {
	Convey("TestRetryGo", t, func() {
		patches := ApplyFuncSeq(DoSomething, []OutputCell{
			{Values: Params{0, errors.New("connection refuse")}, Times: 5},
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
			retry.OnRetry(func(n uint, err error) {
				fmt.Printf("[%v] retry: [%v], err: [%v]\n", time.Now().Format(time.RFC3339), n, err)
			}), // 在重试间隔中执行
			retry.Delay(time.Second),            // 重试间隔 1s
			retry.MaxDelay(3*time.Second),       // 最大重试间隔 3s
			retry.DelayType(retry.BackOffDelay), // 重试策略，时间间隔指数增长
			// retry.DelayType(retry.CombineDelay(retry.BackOffDelay, retry.RandomDelay)), // 指数延迟 + 随机延迟
			// retry.DelayType(retry.CombineDelay(retry.FixedDelay, retry.RandomDelay)), // 固定延迟 + 随机延迟
			// retry.MaxJitter(500*time.Millisecond), // 最大随机延迟
			retry.Attempts(6),         // 最大重试次数 6
			retry.LastErrorOnly(true), // 仅返回最后一次错误
			retry.RetryIf(func(err error) bool {
				return err.Error() == "connection refuse"
			}), // 仅当满足条件时才执行重试
		)

		So(err, ShouldBeNil)
		So(res, ShouldEqual, 123)
	})
}
