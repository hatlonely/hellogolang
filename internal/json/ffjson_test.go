package json

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/pquerna/ffjson/ffjson"
)

// 需要先生成 ffjson 需要的文件
// ffjson jbook.go

func TestFfjson(t *testing.T) {
	Convey("Given 一本书的定义", t, func() {
		Convey("When 序列化", func() {
			book := JBook{
				BookId: 12125924,
				Title:  "人类简史-从动物到上帝",
				Author: "尤瓦尔·赫拉利",
				Price:  40.8,
				Hot:    true,
				Weight: 100,
			}

			data, err := ffjson.Marshal(&book)
			So(err, ShouldBeNil)

			Convey("Then 序列化的结果正确", func() {
				So(string(data), ShouldEqual, `{"id":12125924,"name":"人类简史-从动物到上帝","author":"尤瓦尔·赫拉利","price":40.8,"hot":"true"}`)
			})
		})

		Convey("When 反序列化", func() {
			var book JBook
			str := `{"id":12125925,"name":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":"true"}`

			err := ffjson.Unmarshal([]byte(str), &book)
			So(err, ShouldBeNil)

			Convey("Then 反序列化的结果正确", func() {
				So(book, ShouldResemble, JBook{
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
