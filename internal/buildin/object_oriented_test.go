package buildin

import (
	"fmt"
	"testing"
)

type Sayer interface {
	Say()
	SayHello()
	SayWorld()
}

type Animal struct {
}

func (a *Animal) Say() {
	fmt.Println("not implement")
}

func (a *Animal) SayHello() {
	a.Say()
	fmt.Println("hello")
}

func SayWorld(sayer Sayer) {
	sayer.Say()
	fmt.Println("world")
}

func (a *Animal) SayWorld() {
	SayWorld(a)
}

type Dog struct {
	Animal
}

func (d *Dog) Say() {
	fmt.Println("wang wang wang")
}

func (d *Dog) SayWorld() {
	SayWorld(d)
}

type Cat struct {
	Animal
}

func (c *Cat) Say() {
	fmt.Println("miao miao miao")
}

func (c *Cat) SayWorld() {
	SayWorld(c)
}

type Mouse struct {
	Animal
}

func TestObjectOriented(t *testing.T) {
	var sayer Sayer

	// 有继承，有限的多态特性，父类的实现中无法调用子类的方法
	sayer = &Dog{}
	sayer.Say()      // wang wang wang
	sayer.SayHello() // not implement hello
	sayer.SayWorld() // wang wang wang world

	fmt.Println()

	sayer = &Cat{}
	sayer.Say()
	sayer.SayHello()
	sayer.SayWorld()

	fmt.Println()

	sayer = &Mouse{}
	sayer.Say()
	sayer.SayHello()
	sayer.SayWorld()

	fmt.Println()
}
