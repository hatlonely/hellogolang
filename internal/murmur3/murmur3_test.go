package murmur3

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spaolacci/murmur3"
)

func TestMurmur3(t *testing.T) {
	Convey("Given 一个字符串", t, func() {
		message := "hello golang"
		Convey("When 使用murmur3哈希", func() {
			val := murmur3.Sum64([]byte(message))
			Convey("Then hash的结果正确", func() {
				So(val, ShouldEqual, uint64(12958155624392757344))
			})
		})
	})
}
