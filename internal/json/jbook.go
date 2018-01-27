package json

type JBook struct {
	BookId int     `json:"id"`
	Title  string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price,omitempty"` // omitempty 表示忽略控制
	Hot	   bool    `json:"hot,string"`		// string 表示输出成字符串
	Weight int     `json:"-"`               // '-' 表示不序列化
}
