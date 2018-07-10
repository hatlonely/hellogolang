package json

// Book book
type Book struct {
	BookID int     `json:"bookID"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
	Hot    bool    `json:"hot"`
	Weight int     `json:"-"`
}

// FBook for ffjson
type FBook Book

// EBook for easyjosn
// easyjson:json
// ffjson:skip
type EBook Book
