package bitset

import (
	"strconv"
)

// NewBitSet 创建一个新的 bitset
func NewBitSet() BitSet {
	return BitSet{
		bit: uint64(0),
	}
}

// BitSet 用一个64位的数来表示一个 set，每一位代表一个值
type BitSet struct {
	bit uint64
}

// Add 添加一个数 i
func (bs *BitSet) Add(i uint64) {
	bs.bit |= 1 << i
}

// Del 删除一个数 i
func (bs *BitSet) Del(i uint64) {
	bs.bit &= ^(1 << i)
}

// Has 是否存在 i
func (bs BitSet) Has(i uint64) bool {
	return bs.bit&(1<<i) != 0
}

// Empty 判断是否为空
func (bs BitSet) Empty() bool {
	return bs.bit == 0
}

// MarshalJSON json 序列化
func (bs BitSet) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatUint(bs.bit, 2) + `"`), nil
}

// UnmarshalJSON json 序列化
func (bs *BitSet) UnmarshalJSON(buf []byte) error {
	var err error
	bs.bit, err = strconv.ParseUint(string(buf[1:len(buf)-1]), 2, 64)
	return err
}
