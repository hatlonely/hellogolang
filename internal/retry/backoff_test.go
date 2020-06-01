package retry_test

import (
	"context"
	"errors"
	"testing"
	"time"

	. "github.com/agiledragon/gomonkey"
	"github.com/cenkalti/backoff/v4"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBackoff1(t *testing.T) {
	Convey("TestBackoff1", t, func() {
		patches := ApplyFuncSeq(DoSomething, []OutputCell{
			{Values: Params{0, errors.New("connection refuse")}, Times: 3},
			{Values: Params{123, nil}},
		})
		defer patches.Reset()

		// 第一次 current = InitialInterval
		// 下次调度时间  random(current * (1- RandomizationFactor), current * (1 + RandomizationFactor)) * Multiplier
		// 退出机制 current >= MaxElapsedTime
		bo := backoff.NewExponentialBackOff()
		bo.InitialInterval = 500 * time.Millisecond
		bo.RandomizationFactor = 0.5
		bo.Multiplier = 1.5
		bo.MaxInterval = 60 * time.Second
		bo.MaxElapsedTime = 15 * time.Minute
		bo.Reset()

		var res int
		err := backoff.Retry(func() error {
			i, err := DoSomething()
			if err != nil {
				return err
			}

			res = i
			return nil
		}, bo)

		So(err, ShouldBeNil)
		So(res, ShouldEqual, 123)
	})
}

func TestBackoff2(t *testing.T) {
	Convey("TestBackoff1", t, func() {
		patches := ApplyFunc(DoSomething, func() (int, error) {
			return 0, errors.New("connection refuse")
		})
		defer patches.Reset()

		// 第一次 current = InitialInterval
		// 下次调度时间  random(current * (1- RandomizationFactor), current * (1 + RandomizationFactor)) * Multiplier
		// 退出机制 current >= MaxElapsedTime
		bo := backoff.NewExponentialBackOff()
		bo.InitialInterval = 20 * time.Millisecond
		bo.RandomizationFactor = 0.5
		bo.Multiplier = 1.5
		bo.MaxInterval = 500 * time.Millisecond
		bo.MaxElapsedTime = 3 * time.Second
		bo.Reset()

		var res int
		err := backoff.Retry(func() error {
			i, err := DoSomething()
			if err != nil {
				return err
			}

			res = i
			return nil
		}, bo)

		So(err, ShouldNotBeNil)
		So(res, ShouldEqual, 0)
	})
}

func TestBackoff3(t *testing.T) {
	Convey("TestBackoff1", t, func() {
		patches := ApplyFuncSeq(DoSomething, []OutputCell{
			{Values: Params{0, errors.New("connection refuse")}, Times: 3},
			{Values: Params{123, nil}},
		})
		defer patches.Reset()

		// 第一次 current = InitialInterval
		// 下次调度时间  random(current * (1- RandomizationFactor), current * (1 + RandomizationFactor)) * Multiplier
		// 退出机制 current >= MaxElapsedTime
		bo := backoff.NewExponentialBackOff()
		bo.InitialInterval = 500 * time.Millisecond
		bo.RandomizationFactor = 0.5
		bo.Multiplier = 1.5
		bo.MaxInterval = 60 * time.Second
		bo.MaxElapsedTime = 15 * time.Minute
		bo.Reset()

		ctx := backoff.WithContext(bo, context.Background())
		var res int
		err := backoff.Retry(func() error {
			i, err := DoSomething()
			if err != nil {
				return err
			}

			res = i
			return nil
		}, ctx)

		So(err, ShouldBeNil)
		So(res, ShouldEqual, 123)
	})
}
