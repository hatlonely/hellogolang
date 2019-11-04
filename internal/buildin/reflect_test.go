package buildin

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type A struct {
	F1 int    `json:"f1"`
	F2 string `json:"f2"`
}

func (a *A) Add(b int) int {
	a.F1 += b
	return a.F1
}

func (a A) Mul(b int) int {
	return a.F1 * b
}

func (a A) Sum(vi ...int) int {
	sum := 0
	for _, i := range vi {
		sum += i
	}

	return sum + a.F1
}

func TestReflect(t *testing.T) {
	Convey("遍历类", t, func() {
		t := reflect.TypeOf(A{})
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fmt.Println(field.Name)
			fmt.Println(field.Type)
			fmt.Println(field.Tag)
		}
	})

	Convey("获取 tag", t, func() {
		So(reflect.TypeOf(A{}).Field(0).Tag.Get("json"), ShouldEqual, "f1")
		So(reflect.TypeOf(A{}).Field(1).Tag.Get("json"), ShouldEqual, "f2")
	})

	Convey("遍历值", t, func() {
		a := A{
			F1: 10,
			F2: "hatlonely",
		}
		v := reflect.ValueOf(a)
		for i := 0; i < v.NumField(); i++ {
			val := v.Field(i)
			fmt.Println(val.Type().Kind())
			switch val.Type().Kind() {
			case reflect.Int:
				fmt.Println(val.Int())
			case reflect.String:
				fmt.Println(val.String())
			}
		}
	})

	Convey("根据字段名获取值", t, func() {
		a := A{
			F1: 10,
			F2: "hatlonely",
		}

		So(reflect.ValueOf(a).FieldByName("F1").Type().Kind(), ShouldEqual, reflect.Int)
		So(reflect.ValueOf(a).FieldByName("F2").Type().Kind(), ShouldEqual, reflect.String)
		So(reflect.ValueOf(a).FieldByName("F1").Int(), ShouldEqual, 10)
		So(reflect.ValueOf(a).FieldByName("F2").String(), ShouldEqual, "hatlonely")
	})
}

func TestReflectFunc(t *testing.T) {
	Convey("获取方法", t, func() {
		a := A{
			F1: 10,
		}

		So(reflect.ValueOf(a).MethodByName("Mul").Call([]reflect.Value{reflect.ValueOf(20)})[0].Int(), ShouldEqual, 200)
		So(a.F1, ShouldEqual, 10)
		So(reflect.ValueOf(&a).MethodByName("Add").Call([]reflect.Value{reflect.ValueOf(20)})[0].Int(), ShouldEqual, 30)
		So(a.F1, ShouldEqual, 30)
		So(reflect.ValueOf(a).MethodByName("Sum").Call([]reflect.Value{
			reflect.ValueOf(30), reflect.ValueOf(40), reflect.ValueOf(50),
		})[0].Int(), ShouldEqual, 150)
		So(reflect.ValueOf(a).MethodByName("Sum").CallSlice([]reflect.Value{
			reflect.ValueOf([]int{30, 40, 50}),
		})[0].Int(), ShouldEqual, 150)
	})
}
