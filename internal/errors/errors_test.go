package errors_test

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func TestErrors(t *testing.T) {
	err1 := errors.New("err1 message")
	err2 := errors.Wrap(err1, "err2 message")
	err3 := errors.Wrap(err2, "err3 message")

	fmt.Println(err3.Error())
	fmt.Printf("%v\n", errors.Cause(err3))
	fmt.Printf("%T\n", errors.Cause(err3))
	fmt.Printf("%+v\n", err3)
}
