package buildin

import (
	"bytes"
	"encoding/ascii85"
	"io"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// ascii85 数据编码(5 个 ascii 字符表示 4 个字节)，pdf 文档的编码格式
func TestAscii85(t *testing.T) {
	Convey("test ascii85 encode/decode", t, func() {
		{
			buf := make([]byte, 32)
			n := ascii85.Encode(buf, []byte("hello world"))
			So(buf[0:n], ShouldResemble, []byte("BOu!rD]j7BEbo7"))
		}
		{
			buf := make([]byte, 32)
			n, _, _ := ascii85.Decode(buf, []byte("BOu!rD]j7BEbo7"), true)
			So(buf[0:n], ShouldResemble, []byte("hello world"))
		}
	})

	Convey("test ascii85 encoder/decoder", t, func() {
		{
			buf := make([]byte, 4)
			reader := ascii85.NewDecoder(bytes.NewReader([]byte("BOu!rD]j7BEbo7")))
			writer := &bytes.Buffer{}
			for n, err := reader.Read(buf); err != io.EOF; n, err = reader.Read(buf) {
				_, _ = writer.Write(buf[0:n])
			}
			So(writer.String(), ShouldEqual, "hello world")
		}
		{
			buffer := &bytes.Buffer{}
			writer := ascii85.NewEncoder(buffer)
			_, _ = writer.Write([]byte("hello"))
			_, _ = writer.Write([]byte(" "))
			_, _ = writer.Write([]byte("world"))
			_ = writer.Close()
			So(buffer.String(), ShouldEqual, "BOu!rD]j7BEbo7")
		}
	})

	Convey("support chinese", t, func() {
		{
			buf := make([]byte, 32)
			n := ascii85.Encode(buf, []byte("你好世界"))
			So(buf[0:n], ShouldResemble, []byte("jLq5JV7ks\"QKONl"))
		}
		{
			buf := make([]byte, 32)
			n, _, _ := ascii85.Decode(buf, []byte("jLq5JV7ks\"QKONl"), true)
			So(buf[0:n], ShouldResemble, []byte("你好世界"))
		}
	})
}
