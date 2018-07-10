package hash

// StringHasherBKDR bkdrHash 算法
type StringHasherBKDR uint64

// NewStringHasherBKDR 创建一个新的 Hasher
func NewStringHasherBKDR() StringHasherBKDR {
	return StringHasherBKDR(0)
}

// AddStr 增加一个字符串
func (bkdr StringHasherBKDR) AddStr(str string) StringHasherBKDR {
	val := uint64(bkdr)
	for i := 0; i < len(str); i++ {
		val = val*131 + uint64(str[i])
	}
	return StringHasherBKDR(val)
}

// AddInt 添加一个 int 值
func (bkdr StringHasherBKDR) AddInt(i uint64) StringHasherBKDR {
	val := uint64(bkdr)
	val = val*131 + i
	return StringHasherBKDR(val)
}

// Val 转成 uint64 的值
func (bkdr StringHasherBKDR) Val() uint64 {
	return uint64(bkdr)
}
