package yaegi

import (
	"fmt"
	"testing"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func TestYaegiExample(t *testing.T) {
	i := interp.New(interp.Options{})

	err := i.Use(stdlib.Symbols)
	if err != nil {
		panic(err)
	}

	_, err = i.Eval(`import "fmt"`)
	if err != nil {
		panic(err)
	}

	_, err = i.Eval(`fmt.Println("Hello Yaegi")`)
	if err != nil {
		panic(err)
	}
}

const src = `
func Add(a int, b int) int {
	return a + b
}
`

func TestYaegiFunc(t *testing.T) {
	i := interp.New(interp.Options{})

	_, err := i.Eval(src)
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("Add")
	if err != nil {
		panic(err)
	}

	add := v.Interface().(func(int, int) int)
	sum := add(12, 34)
	fmt.Println(sum)
}

func localFunc(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return localFunc(n-1) + localFunc(n-2)
}

// goos: darwin
// goarch: amd64
// pkg: github.com/hatlonely/hellogolang/internal/embed/yaegi
// cpu: Intel(R) Core(TM) i5-6600 CPU @ 3.30GHz
// BenchmarkYaegiFunc
// BenchmarkYaegiFunc/local
// BenchmarkYaegiFunc/local-4         	      21	   53113709 ns/op
// BenchmarkYaegiFunc/yaegi
// BenchmarkYaegiFunc/yaegi-4         	       1	18070525993 ns/op
// PASS
func BenchmarkYaegiFunc(b *testing.B) {
	i := interp.New(interp.Options{})
	v, err := i.Eval(`
func Fab(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fab(n-1) + Fab(n-2)
}
`)
	if err != nil {
		panic(err)
	}
	v, err = i.Eval("Fab")
	yaegiFunc := v.Interface().(func(int) int)

	fmt.Println(localFunc(35))
	fmt.Println(yaegiFunc(35))

	b.Run("local", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			localFunc(35)
		}
	})

	b.Run("yaegi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			yaegiFunc(35)
		}
	})
}
