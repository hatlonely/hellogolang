package json

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
 	"github.com/json-iterator/go"
)

// 只需要把
// import "encoding/json"
//
// 替换成下面两句即可，完全兼容
// import "github.com/json-iterator/go"
// var json = jsoniter.ConfigCompatibleWithStandardLibrary

var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary

func TestJsonIterator(t *testing.T) {
	Convey("Given 一本书的定义", t, func() {
		type Book struct {
			BookId int     `json:"id"`
			Title  string  `json:"name"`
			Author string  `json:"author"`
			Price  float64 `json:"price,omitempty"` // omitempty 表示忽略控制
			Hot	   bool    `json:"hot,string"`		// string 表示输出成字符串
			Weight int     `json:"-"`               // '-' 表示不序列化
		}

		Convey("When 序列化", func() {
			book := Book{
				BookId: 12125924,
				Title:  "人类简史-从动物到上帝",
				Author: "尤瓦尔·赫拉利",
				Price:  40.8,
				Hot: true,
				Weight: 100,
			}

			data, err := jsonIterator.Marshal(&book)
			So(err, ShouldBeNil)

			Convey("Then 序列化的结果正确", func() {
				So(string(data), ShouldEqual, `{"id":12125924,"name":"人类简史-从动物到上帝","author":"尤瓦尔·赫拉利","price":40.8,"hot":"true"}`)
			})
		})

		Convey("When 反序列化", func() {
			var book Book
			str := `{"id":12125925,"name":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":"true"}`

			err := jsonIterator.Unmarshal([]byte(str), &book)
			So(err, ShouldBeNil)

			Convey("Then 反序列化的结果正确", func() {
				So(book, ShouldResemble, Book{
					BookId: 12125925,
					Title:  "未来简史-从智人到智神",
					Author: "尤瓦尔·赫拉利",
					Price:  40.8,
					Hot: true,
					Weight: 0,
				})
			})
		})
	})
}
