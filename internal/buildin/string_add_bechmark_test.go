package buildin

import (
	"fmt"
	"testing"
	"strings"
	"bytes"
	"strconv"
)

// 下面这个场景关于字符串连接集中方式的性能测试对比
// 关于这个问题的讨论，详见：

func BenchmarkAddStringWithOperator(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = hello + "," + world
	}
}

func BenchmarkAddStringWithSprintf(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s,%s", hello, world)
	}
}

func BenchmarkAddStringWithJoin(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{hello, world}, ",")
	}
}

func BenchmarkAddStringWithBuffer(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < 1000; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(hello)
		buffer.WriteString(",")
		buffer.WriteString(world)
		_ = buffer.String()
	}
}

func BenchmarkAddStringNumberWithOperator(b *testing.B) {
	hello := "hello"
	world := int64(1234567890)
	for i := 0; i < b.N; i++ {
		_ = hello + "," + strconv.FormatInt(world, 10)
	}
}

func BenchmarkAddStringNumberWithSprintf(b *testing.B) {
	hello := "hello"
	world := int64(1234567890)
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s,%d", hello, world)
	}
}

func BenchmarkAddStringNumberWithBuffer(b *testing.B) {
	hello := "hello"
	world := int64(1234567890)
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(hello)
		buffer.WriteString(",")
		buffer.WriteString(strconv.FormatInt(world, 10))
		_ = buffer.String()
	}
}
