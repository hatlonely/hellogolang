package hash

// StringHasherDJB DJBHash 算法
type StringHasherDJB uint64

// NewStringHasherDJB 创建一个新的 Hasher
func NewStringHasherDJB() StringHasherDJB {
	return StringHasherDJB(5381)
}

// AddStr 增加一个字符串
func (djb StringHasherDJB) AddStr(str string) StringHasherDJB {
	val := uint64(djb)
	for i := 0; i < len(str); i++ {
		val = ((val << 5) + val) + uint64(str[i])
	}
	return StringHasherDJB(val)
}

// AddInt 添加一个 int 值
func (djb StringHasherDJB) AddInt(i uint64) StringHasherDJB {
	val := uint64(djb)
	val = ((val << 5) + val) + i
	return StringHasherDJB(val)
}

// Val 转成 uint64 的值
func (djb StringHasherDJB) Val() uint64 {
	return uint64(djb)
}
