package json

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/ugorji/go/codec"
)

// 需要先下载工具
// go get -u github.com/ugorji/go/codec/codecgen

func TestCodecJson(t *testing.T) {
	Convey("test codec json", t, func() {
		type Book struct {
			BookId int64   `json:"id"`
			Title  string  `json:"title"`
			Author string  `json:"author"`
			Price  float64 `json:"price"`
			Hot    bool    `json:"hot"`
			Weight int64   `json:"-"`
		}

		Convey("marshal", func() {
			book := Book{
				BookId: 12125924,
				Title:  "人类简史-从动物到上帝",
				Author: "尤瓦尔·赫拉利",
				Price:  40.8,
				Hot:    true,
				Weight: 100,
			}

			buf := bytes.NewBuffer(make([]byte, 0, 64))
			writer := bufio.NewWriter(buf)
			encoder := codec.NewEncoder(writer, &codec.JsonHandle{})
			So(encoder.Encode(&book), ShouldBeNil)
			_ = writer.Flush()
			So(len(buf.String()), ShouldEqual, len(`{"author":"尤瓦尔·赫拉利","hot":true,"id":12125924,"price":40.8,"title":"人类简史-从动物到上帝"}`))
			//So(buf.String(), ShouldEqual, `{"author":"尤瓦尔·赫拉利","hot":true,"id":12125924,"price":40.8,"title":"人类简史-从动物到上帝"}`)
		})

		Convey("unmarshal", func() {
			var book Book
			str := `{"id":12125925,"title":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`

			reader := bufio.NewReader(strings.NewReader(str))
			decoder := codec.NewDecoder(reader, &codec.JsonHandle{})
			So(decoder.Decode(&book), ShouldBeNil)
			So(book, ShouldResemble, Book{
				BookId: 12125925,
				Title:  "未来简史-从智人到智神",
				Author: "尤瓦尔·赫拉利",
				Price:  40.8,
				Hot:    true,
				Weight: 0,
			})
		})
	})

}
