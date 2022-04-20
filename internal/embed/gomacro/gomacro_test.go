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

func BenchmarkNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fab(35)
	}
}

func BenchmarkGoMacro(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		interp.Eval(`
fab(35)
`)
	}
}
