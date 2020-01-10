package buildin

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// 可变长度的 int 型，根据数值的大小选择不同的长度编码，可节省空间
func TestVarInt(t *testing.T) {
	Convey("test binary", t, func() {
		{
			buf := make([]byte, binary.MaxVarintLen64)
			n := binary.PutUvarint(buf, 12345678901234567890)
			So(n, ShouldEqual, 10)
			So(hex.EncodeToString(buf[0:n]), ShouldEqual, "d295fcd8ceb1aaaaab01")

			i, n := binary.Uvarint(buf)
			So(n, ShouldEqual, 10)
			So(i, ShouldEqual, uint64(12345678901234567890))
		}
		{
			buf := make([]byte, binary.MaxVarintLen64)
			n := binary.PutUvarint(buf, 123)
			So(n, ShouldEqual, 1)
			So(hex.EncodeToString(buf[0:n]), ShouldEqual, "7b")

			i, n := binary.Uvarint(buf)
			So(n, ShouldEqual, 1)
			So(i, ShouldEqual, 123)
		}
		{
			buf := make([]byte, binary.MaxVarintLen64)
			n := binary.PutVarint(buf, 123)
			So(n, ShouldEqual, 2)
			So(hex.EncodeToString(buf[0:n]), ShouldEqual, "f601")

			i, n := binary.Varint(buf)
			So(n, ShouldEqual, 2)
			So(i, ShouldEqual, 123)
		}
	})
}

func TestByteOrder(t *testing.T) {
	Convey("test byte order", t, func() {
		{
			buf := make([]byte, 2)

			binary.BigEndian.PutUint16(buf, 1234)
			So(hex.EncodeToString(buf), ShouldEqual, "04d2")
			So(binary.BigEndian.Uint16(buf), ShouldEqual, 1234)

			binary.LittleEndian.PutUint16(buf, 1234)
			So(hex.EncodeToString(buf), ShouldEqual, "d204")
			So(binary.LittleEndian.Uint16(buf), ShouldEqual, 1234)
		}
		{
			buf := make([]byte, 4)

			binary.BigEndian.PutUint32(buf, 12341234)
			So(hex.EncodeToString(buf), ShouldEqual, "00bc4ff2")
			So(binary.BigEndian.Uint32(buf), ShouldEqual, 12341234)

			binary.LittleEndian.PutUint32(buf, 12341234)
			So(hex.EncodeToString(buf), ShouldEqual, "f24fbc00")
			So(binary.LittleEndian.Uint32(buf), ShouldEqual, 12341234)
		}
		{
			buf := make([]byte, 8)

			binary.BigEndian.PutUint64(buf, 12345678901234567890)
			So(hex.EncodeToString(buf), ShouldEqual, "ab54a98ceb1f0ad2")
			So(binary.BigEndian.Uint64(buf), ShouldEqual, uint64(12345678901234567890))

			binary.LittleEndian.PutUint64(buf, 12345678901234567890)
			So(hex.EncodeToString(buf), ShouldEqual, "d20a1feb8ca954ab")
			So(binary.LittleEndian.Uint64(buf), ShouldEqual, uint64(12345678901234567890))
		}
	})
}
