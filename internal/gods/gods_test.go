package gods

import (
	"testing"

	"github.com/emirpasic/gods/lists/singlylinkedlist"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSinglylinkedlist(t *testing.T) {
	Convey("单链表测试", t, func() {
		li1 := singlylinkedlist.New()
		for i := 0; i < 10; i++ {
			li1.Add(i)
		}
		So(li1.Size(), ShouldEqual, 10)
		li2 := li1.Select(func(idx int, value interface{}) bool {
			return value.(int)%2 == 1
		})
		li2.Each(func(idx int, value interface{}) {
			So(value.(int), ShouldEqual, idx*2+1)
		})
	})
}
