package slice

import "testing"

var N = 300000

func BenchmarkSlice(b *testing.B) {
	b.Run("cap = 0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var arr []int
			for i := 0; i < N; i++ {
				arr = append(arr, i)
			}
		}
	})
	b.Run("cap = 0.1 * len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			arr := make([]int, 0, N/10)
			for i := 0; i < N; i++ {
				arr = append(arr, i)
			}
		}
	})
	b.Run("cap = 0.5 * len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			arr := make([]int, 0, N/2)
			for i := 0; i < N; i++ {
				arr = append(arr, i)
			}
		}
	})
	b.Run("cap = len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			arr := make([]int, 0, N)
			for i := 0; i < N; i++ {
				arr = append(arr, i)
			}
		}
	})
	b.Run("cap = 10 * len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			arr := make([]int, 0, 10*N)
			for i := 0; i < N; i++ {
				arr = append(arr, i)
			}
		}
	})
	b.Run("reuse", func(b *testing.B) {
		var arr []int
		for i := 0; i < b.N; i++ {
			arr = arr[:0]
			for i := 0; i < N; i++ {
				arr = append(arr, i)
			}
		}
	})
}
