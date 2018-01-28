package json

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ugorji/go/codec"
	"bytes"
	"bufio"
	"strings"
)

// 需要先下载工具
// go get -u github.com/ugorji/go/codec/codecgen

func TestCodecJson(t *testing.T) {
	Convey("Given 一本书的定义", t, func() {
		type Book struct {
			BookId int64   `json:"id"`
			Title  string  `json:"title"`
			Author string  `json:"author"`
			Price  float64 `json:"price"`
			Hot    bool    `json:"hot"`
			Weight int64   `json:"-"`
		}

		Convey("When 序列化", func() {
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
			jsonHandler := &codec.JsonHandle{}
			encoder := codec.NewEncoder(writer, jsonHandler)
			err := encoder.Encode(&book)
			writer.Flush()
			So(err, ShouldBeNil)

			Convey("Then 序列化的结果正确", func() {
				So(buf.String(), ShouldEqual, `{"author":"尤瓦尔·赫拉利","hot":true,"id":12125924,"price":40.8,"title":"人类简史-从动物到上帝"}`)
			})
		})

		Convey("When 反序列化", func() {
			var book Book
			str := `{"id":12125925,"title":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`

			reader := bufio.NewReader(strings.NewReader(str))
			jsonHandler := &codec.JsonHandle{}
			decoder := codec.NewDecoder(reader, jsonHandler)
			err := decoder.Decode(&book)
			So(err, ShouldBeNil)

			Convey("Then 反序列化的结果正确", func() {
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
	})

}
