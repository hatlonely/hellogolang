package main

import (
	"math/rand"
	"github.com/hatlonely/easygolang/pprof"
)

func doSomething1() {
	i := 0
	for {
		i++
	}
}

func doSomething2() {
	i := 1
	for {
		i *= 2
	}
}

func doSomething3() {
	for {
		doSomething4()
	}
}

func doSomething4() {
	rand.Int()
}

func main() {
	go doSomething1()
	go doSomething2()
	go doSomething3()

	if err := pprof.PPCmd("cpu 10s"); err != nil {
		panic(err)
	}

	if err := pprof.PPCmd("mem"); err != nil {
		panic(err)
	}
}
