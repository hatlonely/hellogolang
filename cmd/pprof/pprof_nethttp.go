package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"math/rand"
)

func doSomething5() {
	i := 0
	for {
		i++
	}
}

func doSomething6() {
	i := 1
	for {
		i *= 2
	}
}

func doSomething7() {
	for {
		doSomething8()
	}
}

func doSomething8() {
	rand.Int()
}

func main() {
	go doSomething5()
	go doSomething6()
	go doSomething7()
	http.ListenAndServe(fmt.Sprintf("localhost:3000"), nil)
}
