package json

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"github.com/pquerna/ffjson/ffjson"
)

var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary

// type Book struct {
// 	BookID int     `json:"bookID"`
// 	Title  string  `json:"title"`
// 	Author string  `json:"author"`
// 	Price  float64 `json:"price"`
// 	Hot    bool    `json:"hot"`
// 	Weight int     `json:"-"`
// }

var book *Book
var data []byte

func init() {
	book = &Book{
		BookID: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}
	data = []byte(`{"bookID":12125925,"title":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
}

func BenchmarkJsonMarshal(b *testing.B) {
	var data []byte
	b.Run("stdjson marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data, _ = json.Marshal(book)
		}
	})

	b.Run("jsoniter marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data, _ = jsonIterator.Marshal(book)
		}
	})

	fbook := &FBook{}
	*fbook = FBook(*book)
	b.Run("ffjson marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data, _ = ffjson.Marshal(fbook)
		}
	})

	ebook := &EBook{}
	*ebook = EBook(*book)
	b.Run("easyjson marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data, _ = easyjson.Marshal(ebook)
		}
	})

	fmt.Println(string(data))
}

func BenchmarkJsonUnMarshal(b *testing.B) {
	book := &Book{}
	fmt.Println(book)

	b.Run("stdjson marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = json.Unmarshal(data, book)
		}
	})

	b.Run("jsoniter marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = jsonIterator.Unmarshal(data, book)
		}
	})

	fbook := &FBook{}
	*fbook = FBook(*book)
	b.Run("ffjson marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ffjson.Unmarshal(data, fbook)
		}
	})

	ebook := &EBook{}
	*ebook = EBook(*book)
	b.Run("easyjson marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = easyjson.Unmarshal(data, ebook)
		}
	})

	fmt.Println(book)
}
