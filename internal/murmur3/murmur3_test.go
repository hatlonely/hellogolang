package murmur3

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spaolacci/murmur3"
)

func TestMurmur3(t *testing.T) {
	Convey("test murmur3", t, func() {
		h := murmur3.New64()
		So(h.Size(), ShouldEqual, 16)
		_, _ = h.Write([]byte("hello world"))
		So(h.Sum64(), ShouldEqual, uint64(5998619086395760910))

		So(murmur3.Sum64([]byte("hello world")), ShouldEqual, uint64(5998619086395760910))
	})
}
