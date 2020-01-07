package buildin

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestList(t *testing.T) {
	listNew := func(vs []interface{}) *list.List {
		li := list.New()
		for _, v := range vs {
			li.PushBack(v)
		}
		return li
	}

	listToSlice := func(l *list.List) []interface{} {
		var vs []interface{}
		for e := l.Front(); e != nil; e = e.Next() {
			vs = append(vs, e.Value)
		}
		return vs
	}

	Convey("test list", t, func() {
		{
			li := listNew([]interface{}{1, 2, 3, 4, 5})
			So(listToSlice(li), ShouldResemble, []interface{}{1, 2, 3, 4, 5})
			So(li.Len(), ShouldEqual, 5)
			So(li.Front().Value, ShouldEqual, 1)
			So(li.Back().Value, ShouldEqual, 5)
		}
		{
			li := listNew([]interface{}{1, 2, 3, 4, 5})
			li.PushFront(0)
			li.PushBack(6)
			So(listToSlice(li), ShouldResemble, []interface{}{0, 1, 2, 3, 4, 5, 6})
		}
		{
			li := listNew([]interface{}{4, 5, 6})
			li.PushFrontList(listNew([]interface{}{1, 2, 3}))
			li.PushBackList(listNew([]interface{}{7, 8, 9}))
			So(listToSlice(li), ShouldResemble, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})
		}
		{
			li := listNew([]interface{}{1, 2, 3, 4, 5})
			li.Remove(li.Front().Next().Next())
			So(listToSlice(li), ShouldResemble, []interface{}{1, 2, 4, 5})
		}
		{
			li := listNew([]interface{}{1, 2, 3, 4, 5})
			li.InsertBefore(0, li.Front())
			li.InsertAfter(6, li.Back())
			So(listToSlice(li), ShouldResemble, []interface{}{0, 1, 2, 3, 4, 5, 6})
		}
	})
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func TestHeap(t *testing.T) {
	Convey("test heap", t, func() {
		h := &IntHeap{2, 1, 5}
		heap.Init(h)
		heap.Push(h, 3)
		for len(*h) > 0 {
			fmt.Println(heap.Pop(h))
		}
	})
}

func TestRing(t *testing.T) {
	Convey("test ring", t, func() {
		r := ring.New(5)
		for i := 0; i < 5; i++ {
			r.Value = i
			r = r.Next()
		}
		for i := 0; i < r.Len(); i++ {
			fmt.Println(r.Value)
			r = r.Next()
		}
		r.Do(func(v interface{}) {
			fmt.Println(v)
		})
	})
}
