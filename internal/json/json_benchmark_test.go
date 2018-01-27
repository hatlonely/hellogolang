package json

import (
	"testing"
	"encoding/json"
	"github.com/json-iterator/go"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/mailru/easyjson"
)

// 运行性能测试
// go test -bench=. *

type Book struct {
	BookId int     `json:"id"`
	Title  string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
	Hot    bool    `json:"hot"`
	Weight int     `json:"-"`
}

func BenchmarkStdJsonMarshal(b *testing.B) {
	book := Book{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	for i := 0; i < b.N; i++ {
		json.Marshal(&book)
	}
}

func BenchmarkJsonIteratorMarshal(b *testing.B) {
	book := Book{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	for i := 0; i < b.N; i++ {
		jsonIterator.Marshal(&book)
	}
}

func BenchmarkFfjsonMarshal(b *testing.B) {
	book := FBook{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	for i := 0; i < b.N; i++ {
		ffjson.Marshal(&book)
	}
}

func BenchmarkEasyjsonMarshal(b *testing.B) {
	book := EBook{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	for i := 0; i < b.N; i++ {
		easyjson.Marshal(&book)
	}
}

func BenchmarkStdJsonUnMarshal(b *testing.B) {
	data := []byte(`{"id":12125924,"name":"人类简史-从动物到上帝","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book Book

	for i := 0; i < b.N; i++ {
		json.Unmarshal(data, &book)
	}
}

func BenchmarkJsonIteratorUnMarshal(b *testing.B) {
	data := []byte(`{"id":12125924,"name":"人类简史-从动物到上帝","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book Book

	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	for i := 0; i < b.N; i++ {
		jsonIterator.Unmarshal(data, &book)
	}
}

func BenchmarkFfjsonUnMarshal(b *testing.B) {
	data := []byte(`{"id":12125924,"name":"人类简史-从动物到上帝","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book FBook

	for i := 0; i < b.N; i++ {
		ffjson.Unmarshal(data, &book)
	}
}

func BenchmarkEasyjsonUnMarshal(b *testing.B) {
	data := []byte(`{"id":12125924,"name":"人类简史-从动物到上帝","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book EBook

	for i := 0; i < b.N; i++ {
		easyjson.Unmarshal(data, &book)
	}
}
