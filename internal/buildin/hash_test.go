package buildin

import (
	"encoding/binary"
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdler32(t *testing.T) {
	Convey("test adler32", t, func() {
		h := adler32.New()
		So(h.Size(), ShouldEqual, 4)
		n, _ := h.Write([]byte("hello world"))
		So(n, ShouldEqual, len("hello world"))
		So(h.Sum32(), ShouldEqual, 436929629)
		So(binary.BigEndian.Uint32(h.Sum(nil)), ShouldEqual, 436929629)

		So(adler32.Checksum([]byte("hello world")), ShouldEqual, 436929629)
	})
}

func TestCrc32(t *testing.T) {
	Convey("test crc32", t, func() {
		h := crc32.NewIEEE()
		So(h.Size(), ShouldEqual, 4)
		n, _ := h.Write([]byte("hello world"))
		So(n, ShouldEqual, len("hello world"))
		So(h.Sum32(), ShouldEqual, 222957957)
		So(binary.BigEndian.Uint32(h.Sum(nil)), ShouldEqual, 222957957)

		So(crc32.ChecksumIEEE([]byte("hello world")), ShouldEqual, 222957957)
	})
}

func TestCrc64(t *testing.T) {
	Convey("test crc64", t, func() {
		h := crc64.New(crc64.MakeTable(crc64.ECMA))
		So(h.Size(), ShouldEqual, 8)
		n, _ := h.Write([]byte("hello world"))
		So(n, ShouldEqual, len("hello world"))
		So(h.Sum64(), ShouldEqual, 5981764153023615706)
		So(binary.BigEndian.Uint64(h.Sum(nil)), ShouldEqual, 5981764153023615706)
	})
}

func TestFnv(t *testing.T) {
	Convey("test fnv32", t, func() {
		h := fnv.New32()
		So(h.Size(), ShouldEqual, 4)
		n, _ := h.Write([]byte("hello world"))
		So(n, ShouldEqual, len("hello world"))
		So(h.Sum32(), ShouldEqual, 1418570095)
		So(binary.BigEndian.Uint32(h.Sum(nil)), ShouldEqual, 1418570095)
	})

	Convey("test fnv64", t, func() {
		h := fnv.New64()
		So(h.Size(), ShouldEqual, 8)
		n, _ := h.Write([]byte("hello world"))
		So(n, ShouldEqual, len("hello world"))
		So(h.Sum64(), ShouldEqual, 9065573210506989167)
		So(binary.BigEndian.Uint64(h.Sum(nil)), ShouldEqual, 9065573210506989167)
	})

	Convey("test fnv128", t, func() {
		h := fnv.New128()
		So(h.Size(), ShouldEqual, 16)
		n, _ := h.Write([]byte("hello world"))
		So(n, ShouldEqual, len("hello world"))
		fmt.Println(len(h.Sum(nil)))
		fmt.Println(binary.BigEndian.Uint64(h.Sum(nil)[0:8]))
		fmt.Println(binary.BigEndian.Uint64(h.Sum(nil)[8:16]))
	})

	Convey("test fnv32a", t, func() {
		h := fnv.New32a()
		So(h.Size(), ShouldEqual, 4)
		n, _ := h.Write([]byte("hello world"))
		So(n, ShouldEqual, len("hello world"))
		So(h.Sum32(), ShouldEqual, 3582672807)
	})
}
