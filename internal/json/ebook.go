package json

type EBook struct {
	BookId int     `json:"id"`
	Title  string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price,omitempty"` // omitempty 表示忽略控制
	Hot    bool    `json:"hot"`             // easyjson，不支持类型指定
	Weight int     `json:"-"`               // '-' 表示不序列化
}
