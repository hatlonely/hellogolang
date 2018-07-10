package bitset

import "testing"

func BenchmarkSetContains(b *testing.B) {
	bitset := NewBitSet()
	hashset := map[uint64]struct{}{}
	for _, i := range []uint64{1, 2, 4, 10} {
		bitset.Add(i)
		hashset[i] = struct{}{}
	}

	b.Run("bitset", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for i := uint64(0); i < uint64(10); i++ {
				_ = bitset.Has(i)
			}
		}
	})

	b.Run("hashset", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for i := uint64(0); i < uint64(10); i++ {
				_, _ = hashset[i]
			}
		}
	})
}
