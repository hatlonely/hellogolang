package json

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"encoding/json"
)

func TestStdJson(t *testing.T) {
	Convey("Given 一本书的定义", t, func() {
		type Book struct {
			BookId int     `json:"id"`
			Title  string  `json:"title"`
			Author string  `json:"author"`
			Price  float64 `json:"price,omitempty"` // omitempty 表示忽略控制
			Hot    bool    `json:"hot,string"`      // string 表示输出成字符串，可选值 [number | boolean | string]
			Weight int     `json:"-"`               // '-' 表示不序列化
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

			data, err := json.Marshal(&book)
			So(err, ShouldBeNil)

			Convey("Then 序列化的结果正确", func() {
				So(string(data), ShouldEqual, `{"id":12125924,"title":"人类简史-从动物到上帝","author":"尤瓦尔·赫拉利","price":40.8,"hot":"true"}`)
			})
		})

		Convey("When 反序列化", func() {
			var book Book
			str := `{"id":12125925,"title":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":"true"}`

			err := json.Unmarshal([]byte(str), &book)
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
