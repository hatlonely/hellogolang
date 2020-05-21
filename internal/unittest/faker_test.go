package unittest

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bxcodec/faker/v3"
	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"
)

type A struct {
	I32 int32
	I64 int64
	Str string

	ID string `faker:"id"`
}

func TestFaker(t *testing.T) {
	Convey("TestFaker", t, func() {
		_ = faker.AddProvider("id", func(v reflect.Value) (interface{}, error) {
			return uuid.NewV4().String(), nil
		})

		a := A{}
		err := faker.FakeData(&a)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v", a)
	})
}
