package unittest_test

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"testing"

	. "github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
)

// -gcflags=all=-l
// go test -cover -coverprofile=coverage.data -gcflags=all=-l ./...
// go tool cover -func=coverage.data -o coverage.txt

func TestGoMonkeyMockFunc(t *testing.T) {
	Convey("TestGoMonkeyMockFunc", t, func() {
		patches := ApplyFunc(rand.Int, func() int {
			return 123
		})
		defer patches.Reset()

		So(rand.Int(), ShouldEqual, 123)
	})
}

func TestGoMonkeyMockFuncSeq(t *testing.T) {
	Convey("TestGoMonkeyMockFuncSeq", t, func() {
		patches := ApplyFuncSeq(rand.Int, []OutputCell{
			{Values: Params{123}},
			{Values: Params{456}, Times: 2},
			{Values: Params{789}},
		})
		defer patches.Reset()

		So(rand.Int(), ShouldEqual, 123)
		So(rand.Int(), ShouldEqual, 456)
		So(rand.Int(), ShouldEqual, 456)
		So(rand.Int(), ShouldEqual, 789)
	})
}

func TestGoMonkeyMockFuncVar(t *testing.T) {
	randInt := rand.Int
	Convey("TestGoMonkeyMockFuncVar", t, func() {
		patches := ApplyFuncVar(&randInt, func() int {
			return 123
		})
		defer patches.Reset()

		So(randInt(), ShouldEqual, 123)
	})
}

func TestGoMonkeyMockFuncVarSeq(t *testing.T) {
	randInt := rand.Int
	Convey("TestGoMonkeyMockFuncVar", t, func() {
		patches := ApplyFuncVarSeq(&randInt, []OutputCell{
			{Values: Params{123}},
			{Values: Params{456}, Times: 2},
			{Values: Params{789}},
		})
		defer patches.Reset()

		So(randInt(), ShouldEqual, 123)
		So(randInt(), ShouldEqual, 456)
		So(randInt(), ShouldEqual, 456)
		So(randInt(), ShouldEqual, 789)
	})
}

func TestGoMonkeyMockVar(t *testing.T) {
	num := 123
	Convey("TestGoMonkeyMockVar", t, func() {
		patches := ApplyGlobalVar(&num, 456)
		defer patches.Reset()

		So(num, ShouldEqual, 456)
	})
}

func TestGoMonkeyMockMethod(t *testing.T) {
	Convey("TestGoMonkeyMockMethod", t, func() {
		patches := ApplyMethod(reflect.TypeOf(&http.Client{}), "Do", func(
			client *http.Client, req *http.Request) (*http.Response, error) {
			return &http.Response{
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("hello world"))),
				StatusCode: http.StatusOK,
				Status:     http.StatusText(http.StatusOK),
			}, nil
		})
		defer patches.Reset()

		client := &http.Client{}
		req, _ := http.NewRequest("GET", "https://www.baidu.com", nil)
		res, err := client.Do(req)

		So(err, ShouldBeNil)
		So(res.Status, ShouldEqual, "OK")
		So(res.StatusCode, ShouldEqual, http.StatusOK)

		body, _ := ioutil.ReadAll(res.Body)

		So(string(body), ShouldEqual, "hello world")
	})
}

func TestGoMonkeyMockMethodSeq(t *testing.T) {
	Convey("TestGoMonkeyMockMethodSeq", t, func() {
		patches := ApplyMethodSeq(reflect.TypeOf(&http.Client{}), "Do", []OutputCell{
			{
				Values: Params{&http.Response{
					Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("hello world 1"))),
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				}, nil},
			},
			{
				Values: Params{&http.Response{
					Body:       ioutil.NopCloser(bytes.NewBuffer([]byte("hello world 2"))),
					StatusCode: http.StatusOK,
					Status:     http.StatusText(http.StatusOK),
				}, nil},
			},
		})
		defer patches.Reset()

		client := &http.Client{}
		req, _ := http.NewRequest("GET", "https://www.baidu.com", nil)

		{
			res, _ := client.Do(req)
			body, _ := ioutil.ReadAll(res.Body)
			So(string(body), ShouldEqual, "hello world 1")
		}
		{
			res, _ := client.Do(req)
			body, _ := ioutil.ReadAll(res.Body)
			So(string(body), ShouldEqual, "hello world 2")
		}
	})
}
