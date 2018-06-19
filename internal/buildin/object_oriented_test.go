package buildin

import (
	"fmt"
	"testing"
)

type Sayer interface {
	Say()
}

type Animal struct {
}

func (a *Animal) Say() {
	fmt.Println("not implement")
}

type Dog struct {
	Animal
}

func (d *Dog) Say() {
	fmt.Println("wang wang wang")
}

type Cat struct {
	Animal
}

func (d *Cat) Say() {
	fmt.Println("miao miao miao")
}

type Mouse struct {
	Animal
}

func TestObjectOriented(t *testing.T) {
	var sayer Sayer

	sayer = &Dog{}
	sayer.Say()

	sayer = &Cat{}
	sayer.Say()

	sayer = &Mouse{}
	sayer.Say()

	t.Error(1)
}
