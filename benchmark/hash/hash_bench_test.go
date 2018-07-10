package hash

import (
	"fmt"
	"hash/fnv"
	"testing"

	"github.com/spaolacci/murmur3"
)

var buf = []byte("0a9599d1-0920-4e97-b8fe-8407ccf0f387")
var str = string(buf)
var str1 = "hello"
var str2 = "world"
var int1 = 1000
var int2 = 987654321

func BenchmarkHashString(b *testing.B) {
	b.Run("murmur3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = murmur3.Sum64(buf)
		}
	})

	b.Run("murmur3 str", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = murmur3.Sum64([]byte(str))
		}
	})

	b.Run("fnv", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := fnv.New64a()
			h.Write(buf)
			_ = h.Sum64()
		}
	})

	b.Run("djb hasher", func(b *testing.B) {
		hasher := NewStringHasherDJB()
		for i := 0; i < b.N; i++ {
			_ = hasher.AddStr(str).Val()
		}
	})

	b.Run("bkdr hasher", func(b *testing.B) {
		hasher := NewStringHasherBKDR()
		for i := 0; i < b.N; i++ {
			_ = hasher.AddStr(str).Val()
		}
	})

	b.Run("bkdr fixed hasher", func(b *testing.B) {
		hasher := NewStringHasherBKDRFixed()
		for i := 0; i < b.N; i++ {
			_ = hasher.Hash(str).Val()
		}
	})
}

func BenchmarkHashMultiStringHash(b *testing.B) {
	b.Run("murmur3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = murmur3.Sum64([]byte(str1 + str2))
		}
	})

	b.Run("fnv", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := fnv.New64a()
			h.Write([]byte(str1))
			h.Write([]byte(str2))
			_ = h.Sum64()
		}
	})

	b.Run("djb hasher", func(b *testing.B) {
		hasher := NewStringHasherDJB()
		for i := 0; i < b.N; i++ {
			_ = hasher.AddStr(str1).AddStr(str2).Val()
		}
	})

	b.Run("bkdr hasher", func(b *testing.B) {
		hasher := NewStringHasherBKDR()
		for i := 0; i < b.N; i++ {
			_ = hasher.AddStr(str1).AddStr(str2).Val()
		}
	})
}

func BenchmarkHashStringAndInt(b *testing.B) {
	b.Run("murmur3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = murmur3.Sum64([]byte(fmt.Sprintf("%s-%s-%d-%d", str1, str2, int1, int2)))
		}
	})

	b.Run("fnv", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := fnv.New64a()
			h.Write([]byte(fmt.Sprintf("%s-%s-%d-%d", str1, str2, int1, int2)))
			_ = h.Sum64()
		}
	})

	b.Run("djb hasher", func(b *testing.B) {
		hasher := NewStringHasherDJB()
		for i := 0; i < b.N; i++ {
			_ = hasher.AddStr(str1).AddStr("-").AddStr(str2).AddStr("-").AddInt(uint64(int1)).AddStr("-").AddInt(uint64(int2)).Val()
		}
	})

	b.Run("bkdr hasher", func(b *testing.B) {
		hasher := NewStringHasherBKDR()
		for i := 0; i < b.N; i++ {
			_ = hasher.AddStr(str1).AddStr("-").AddStr(str2).AddStr("-").AddInt(uint64(int1)).AddStr("-").AddInt(uint64(int2)).Val()
		}
	})
}
