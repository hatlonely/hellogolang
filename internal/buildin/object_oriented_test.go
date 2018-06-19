package buildin

import (
	"fmt"
	"testing"
)

type Sayer interface {
	Say(message string)
	SayHi()
}

func SayHi(s Sayer) {
	s.Say("Hi")
}

type Animal struct {
	Name string
}

func (a *Animal) Say(message string) {
	fmt.Printf("Animal[%v] say: %v\n", a.Name, message)
}

func (a *Animal) SayHi() {
	SayHi(a)
}

type Dog struct {
	Animal
}

func (d *Dog) Say(message string) {
	fmt.Printf("Dog[%v] say: %v\n", d.Name, message)
}

type Cat struct {
	Animal
}

func (c *Cat) Say(message string) {
	fmt.Printf("Cat[%v] say: %v\n", c.Name, message)
}

func (c *Cat) SayHi() {
	SayHi(c)
}

func TestObjectOriented(t *testing.T) {
	var sayer Sayer

	sayer = &Dog{Animal{Name: "Yoda"}}
	sayer.Say("hello world") // Dog[Yoda] say: hello world
	sayer.SayHi()            // Animal[Yoda] say: Hi

	sayer = &Cat{Animal{Name: "Jerry"}}
	sayer.Say("hello world") // Cat[Jerry] say: hello world
	sayer.SayHi()            // Cat[Jerry] say: Hi
}
