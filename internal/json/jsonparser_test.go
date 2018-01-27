package json

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/buger/jsonparser"
)

func TestJsonparser(t *testing.T) {
	Convey("Given 一本书的定义", t, func() {
		type Book struct {
			BookId int64   `json:"id"`
			Title  string  `json:"title"`
			Author string  `json:"author"`
			Price  float64 `json:"price"`
			Hot    bool    `json:"hot"`
			Weight int64   `json:"-"`
		}

		Convey("When 反序列化", func() {
			str := `{"id":12125925,"title":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`
			buf := []byte(str)

			var err error
			var book Book
			book.BookId, err = jsonparser.GetInt(buf, "id")
			So(err, ShouldBeNil)
			book.Title, err = jsonparser.GetString(buf, "title")
			So(err, ShouldBeNil)
			book.Author, err = jsonparser.GetString(buf, "author")
			So(err, ShouldBeNil)
			book.Price, err = jsonparser.GetFloat(buf, "price")
			So(err, ShouldBeNil)
			book.Hot, err = jsonparser.GetBoolean(buf, "hot")

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
