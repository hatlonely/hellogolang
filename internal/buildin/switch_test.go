package buildin

import (
	"fmt"
	"testing"
)

func TestSwitchCase(t *testing.T) {
	i := 1
	switch i {
	case 1:
		fmt.Println(1)
	case 2, 3, 4:
		fmt.Println(i)
	default:
		fmt.Println(100)
	}
}

func BenchmarkSwitch(b *testing.B) {
	b.Run("benchmark switch", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			switch i {
			case 1:
			case 2, 3, 4:
			default:
			}
		}
	})

	b.Run("benchmark if", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if i == 1 {
			} else if i == 2 || i == 3 || i == 4 {
			} else {
			}
		}
	})
}
