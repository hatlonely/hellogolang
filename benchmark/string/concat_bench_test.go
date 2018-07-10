package string

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var str1 = "hello"
var str2 = "world"
var str3 = "!"

func BenchmarkStringConcat(b *testing.B) {
	b.Run("operator +", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = str1 + str2 + str3 + str1 + str2 + str3 + str1 + str2 + str3
		}
	})
	b.Run("operator +=", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var str string
			str += str1
			str += str2
			str += str3
			str += str1
			str += str2
			str += str3
			str += str1
			str += str2
			str += str3
			_ = str
		}
	})
	b.Run("fmt.Sprintf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%s%s%s%s%s%s%s%s%s", str1, str2, str3, str1, str2, str3, str1, str2, str3)
		}
	})
	b.Run("bytes.Buffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			buf.WriteString(str1)
			buf.WriteString(str2)
			buf.WriteString(str3)
			buf.WriteString(str1)
			buf.WriteString(str2)
			buf.WriteString(str3)
			buf.WriteString(str1)
			buf.WriteString(str2)
			buf.WriteString(str3)
			_ = buf.String()
		}
	})
	b.Run("strings.Join", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = strings.Join([]string{str1, str2, str3, str1, str2, str3, str1, str2, str3}, "")
		}
	})
}
