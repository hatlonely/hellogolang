package gomacro

import (
	"fmt"
	"testing"

	"github.com/cosmos72/gomacro/fast"
)

func TestGoMacro(t *testing.T) {
	interp := fast.New()
	vals, _ := interp.Eval(`1+1`)
	fmt.Println(vals[0].ReflectValue())

	interp.Eval(`
func fab(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}
`)
	vals, _ = interp.Eval(`fab(35)`)

	fmt.Println(vals[0].ReflectValue())
}

func fab(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}

// goos: darwin
// goarch: amd64
// pkg: github.com/hatlonely/hellogolang/internal/embed/gomacro
// cpu: Intel(R) Core(TM) i5-6600 CPU @ 3.30GHz
// BenchmarkGomacro
// 9227465
// 9227465
// BenchmarkGomacro/localFunc
// BenchmarkGomacro/localFunc-4         	      21	  53291006 ns/op
// BenchmarkGomacro/gomacroFunc
// BenchmarkGomacro/gomacroFunc-4       	       1	2303495653 ns/op
// PASS
func BenchmarkGomacro(b *testing.B) {
	interp := fast.New()
	interp.Eval(`
func fab(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}
`)
	fmt.Println(fab(35))
	vals, _ := interp.Eval(`fab(35)`)
	fmt.Println(vals[0].ReflectValue())

	b.Run("localFunc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fab(35)
		}
	})

	b.Run("gomacroFunc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			vals, _ := interp.Eval("fab(35)")
			_ = vals[0].ReflectValue()
		}
	})
}
