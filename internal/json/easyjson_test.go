package json

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/mailru/easyjson"
)

// 需要先生成 easyjson 需要的文件
// easyjson -all ebook.go

func TestEasyjson(t *testing.T) {
	Convey("Given 一本书的定义", t, func() {
		Convey("When 序列化", func() {
			book := EBook{
				BookId: 12125924,
				Title:  "人类简史-从动物到上帝",
				Author: "尤瓦尔·赫拉利",
				Price:  40.8,
				Hot:    true,
				Weight: 100,
			}

			data, err := easyjson.Marshal(&book)
			So(err, ShouldBeNil)

			Convey("Then 序列化的结果正确", func() {
				So(string(data), ShouldEqual, `{"id":12125924,"title":"人类简史-从动物到上帝","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
			})
		})

		Convey("When 反序列化", func() {
			var book EBook
			str := `{"id":12125925,"title":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`

			err := easyjson.Unmarshal([]byte(str), &book)
			So(err, ShouldBeNil)

			Convey("Then 反序列化的结果正确", func() {
				So(book, ShouldResemble, EBook{
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