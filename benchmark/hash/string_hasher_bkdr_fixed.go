package hash

import (
	"reflect"
	"unsafe"
)

// StringHasherBKDRFixed bkdrHash 算法
type StringHasherBKDRFixed struct {
	val uint64
}

// NewStringHasherBKDRFixed 创建一个新的 Hasher
func NewStringHasherBKDRFixed() StringHasherBKDRFixed {
	return StringHasherBKDRFixed{val: 0}
}

// Hash 增加一个字符串
func (bkdr StringHasherBKDRFixed) Hash(str string) StringHasherBKDRFixed {
	header := (*reflect.StringHeader)(unsafe.Pointer(&str))
	data := header.Data
	l := len(str)
	t := l / 8

	for i := 0; i < t; i++ {
		bkdr.val = bkdr.val*131 + *(*uint64)(unsafe.Pointer(data + uintptr(i)*unsafe.Sizeof(uint64(0))))
	}
	for i := 8 * t; i < l; i++ {
		bkdr.val = bkdr.val*131 + uint64(str[i])
	}

	return bkdr
}

// Val 转成 uint64 的值
func (bkdr StringHasherBKDRFixed) Val() uint64 {
	return bkdr.val
}
