package buildin

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func TestAdd(t *testing.T) {
	Convey("可变参数", t, func() {
		So(add(1, 2), ShouldEqual, 3)
		So(add(1, 2, 4), ShouldEqual, 7)
	})
}

func MyFunc1(requiredStr string, str1 string, str2 string, int1 int, int2 int) {
	fmt.Println(requiredStr, str1, str2, int1, int2)
}

func TestMyFunc1(t *testing.T) {
	MyFunc1("requiredStr", "defaultStr1", "defaultStr2", 1, 2)
}

type MyFuncOptions struct {
	optionStr1 string
	optionStr2 string
	optionInt1 int
	optionInt2 int
}

var defaultMyFuncOptions = MyFuncOptions{
	optionStr1: "defaultStr1",
	optionStr2: "defaultStr2",
	optionInt1: 1,
	optionInt2: 2,
}

type MyFuncOption func(options *MyFuncOptions)

func WithOptionStr1(str1 string) MyFuncOption {
	return func(options *MyFuncOptions) {
		options.optionStr1 = str1
	}
}

func WithOptionInt1(int1 int) MyFuncOption {
	return func(options *MyFuncOptions) {
		options.optionInt1 = int1
	}
}

func WithOptionStr2AndInt2(str2 string, int2 int) MyFuncOption {
	return func(options *MyFuncOptions) {
		options.optionStr2 = str2
		options.optionInt2 = int2
	}
}

func MyFunc2(requiredStr string, opts ...MyFuncOption) {
	options := defaultMyFuncOptions
	for _, o := range opts {
		o(&options)
	}

	fmt.Println(requiredStr, options.optionStr1, options.optionStr2, options.optionInt1, options.optionInt2)
}

func TestMyFunc2(t *testing.T) {
	MyFunc2("requiredStr")
	MyFunc2("requiredStr", WithOptionStr1("mystr1"))
	MyFunc2("requiredStr", WithOptionStr2AndInt2("mystr2", 22), WithOptionInt1(11))
}
