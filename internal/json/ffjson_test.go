package json

import (
	"testing"

	"github.com/pquerna/ffjson/ffjson"
	. "github.com/smartystreets/goconvey/convey"
)

// 需要先生成 ffjson 需要的文件
// ffjson jbook.go

func TestFfjson(t *testing.T) {
	Convey("test ffjson", t, func() {
		Convey("marshal", func() {
			book := FBook{
				BookId: 12125924,
				Title:  "人类简史-从动物到上帝",
				Author: "尤瓦尔·赫拉利",
				Price:  40.8,
				Hot:    true,
				Weight: 100,
			}

			data, err := ffjson.Marshal(&book)
			So(err, ShouldBeNil)
			So(string(data), ShouldEqual, `{"id":12125924,"title":"人类简史-从动物到上帝","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
		})

		Convey("unmarshal", func() {
			var book FBook
			str := `{"id":12125925,"title":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`

			So(ffjson.Unmarshal([]byte(str), &book), ShouldBeNil)
			So(book, ShouldResemble, FBook{
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
