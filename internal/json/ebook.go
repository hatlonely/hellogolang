package json

type EBook struct {
	BookId int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
	Hot    bool    `json:"hot"`
	Weight int     `json:"-"`
}
