package unittest

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v5"
	. "github.com/smartystreets/goconvey/convey"
)

type Foo struct {
	Bar      string
	Int      int
	Pointer  *int
	Name     string  `fake:"{firstname}"`  // Any available function all lowercase
	Sentence string  `fake:"{sentence:3}"` // Can call with parameters
	RandStr  string  `fake:"{randomstring:[hello,world]}"`
	Skip     *string `fake:"skip"` // Set to "skip" to not generate data for
}

func TestGoFakeIt(t *testing.T) {
	Convey("", t, func() {
		gofakeit.Seed(0)

		Convey("normal", func() {
			fmt.Println(gofakeit.Name())
			fmt.Println(gofakeit.Email())
			fmt.Println(gofakeit.Phone())
			fmt.Println(gofakeit.BS())
			fmt.Println(gofakeit.BeerName())
			fmt.Println(gofakeit.Color())
			fmt.Println(gofakeit.Company())
			fmt.Println(gofakeit.CreditCardNumber(nil))
			fmt.Println(gofakeit.HackerPhrase())
			fmt.Println(gofakeit.JobTitle())
			fmt.Println(gofakeit.Country())
		})

		Convey("struct", func() {
			a := Foo{}
			gofakeit.Struct(&a)
			fmt.Printf("%+v\n", a)
		})

	})

}
