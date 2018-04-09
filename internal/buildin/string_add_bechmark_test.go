package buildin

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// 下面这个场景关于字符串连接集中方式的性能测试对比
// 关于这个问题的讨论，详见：<http://www.hatlonely.com/2018/01/24/golang-%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%9A%84%E5%87%A0%E7%A7%8D%E8%BF%9E%E6%8E%A5%E6%96%B9%E5%BC%8F/>

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
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(hello)
		buffer.WriteString(",")
		buffer.WriteString(world)
		_ = buffer.String()
	}
}

func BenchmarkAddMoreStringWithOperator(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < 100; j++ {
			str += hello + "," + world
		}
	}
}

func BenchmarkAddMoreStringWithBuffer(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		for j := 0; j < 100; j++ {
			buffer.WriteString(hello)
			buffer.WriteString(",")
			buffer.WriteString(world)
		}
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
