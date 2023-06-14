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

func TestReflectType(t *testing.T) {
	Convey("遍历类", t, func() {
		t := reflect.TypeOf(A{})
		So(t.NumField(), ShouldEqual, 2)
		So(t.Field(0).Name, ShouldEqual, "F1")
		So(t.Field(0).Type.Kind(), ShouldEqual, reflect.Int)
		So(t.Field(0).Tag, ShouldEqual, `json:"f1"`)
		So(t.Field(1).Name, ShouldEqual, "F2")
		So(t.Field(1).Type.Kind(), ShouldEqual, reflect.String)
		So(t.Field(1).Tag, ShouldEqual, `json:"f2"`)
	})

	Convey("获取 tag", t, func() {
		So(reflect.TypeOf(A{}).Field(0).Tag.Get("json"), ShouldEqual, "f1")
		So(reflect.TypeOf(A{}).Field(1).Tag.Get("json"), ShouldEqual, "f2")
	})
}

func TestReflectValue(t *testing.T) {
	Convey("遍历值", t, func() {
		a := A{
			F1: 10,
			F2: "hatlonely",
		}
		v := reflect.ValueOf(a)
		So(v.NumField(), ShouldEqual, 2)
		So(v.Field(0).Int(), ShouldEqual, 10)
		So(v.Field(1).String(), ShouldEqual, "hatlonely")
		So(v.Field(0).Type().Kind(), ShouldEqual, reflect.Int)
		So(v.Field(1).Type().Kind(), ShouldEqual, reflect.String)
		So(v.FieldByName("F1").Int(), ShouldEqual, 10)
		So(v.FieldByName("F2").Interface().(string), ShouldEqual, "hatlonely")
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

func interfaceToStruct(d interface{}, v interface{}) error {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("invalid value type")
	}
	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem()

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		switch rt.Field(i).Type.Kind() {
		case reflect.Int:
			field.Set(reflect.ValueOf(10))
		case reflect.String:
			field.Set(reflect.ValueOf("hatlonely"))
		}
	}

	return nil
}

func TestInterfaceToStruct(t *testing.T) {
	Convey("interface to struct", t, func() {
		Convey("case1", func() {
			v := &A{}
			d := map[string]interface{}{
				"f1": 10,
				"f2": "hatlonely",
			}

			So(interfaceToStruct(d, v), ShouldBeNil)
			So(v.F1, ShouldEqual, 10)
			So(v.F2, ShouldEqual, "hatlonely")
		})
	})
}
