package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"math/rand"
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
	http.ListenAndServe(fmt.Sprintf("localhost:3000"), nil)
}
