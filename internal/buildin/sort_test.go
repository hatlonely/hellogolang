package buildin

import (
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type DescIntSlice []int

func (d DescIntSlice) Len() int {
	return len([]int(d))
}

func (d DescIntSlice) Less(i, j int) bool {
	s := []int(d)
	return s[i] > s[j]
}

func (d DescIntSlice) Swap(i, j int) {
	s := []int(d)
	s[i], s[j] = s[j], s[i]
}

func TestSort(t *testing.T) {
	Convey("test sort int string float", t, func() {
		ints := []int{2, 4, 1, 7, 9, 3, 5, 8, 6}
		So(sort.IntsAreSorted(ints), ShouldBeFalse)
		sort.Ints(ints)
		So(sort.IntsAreSorted(ints), ShouldBeTrue)
		So(ints, ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		So(sort.SearchInts(ints, 4), ShouldEqual, 3)

		strs := []string{"key2", "key4", "key1", "key3", "key0"}
		So(sort.StringsAreSorted(strs), ShouldBeFalse)
		sort.Strings(strs)
		So(sort.StringsAreSorted(strs), ShouldBeTrue)
		So(strs, ShouldResemble, []string{"key0", "key1", "key2", "key3", "key4"})
		So(sort.SearchStrings(strs, "key11"), ShouldEqual, 2)

		floats := []float64{7.8, 3.4, 1.2, 9.0, 5.6}
		So(sort.Float64sAreSorted(floats), ShouldBeFalse)
		sort.Float64s(floats)
		So(sort.Float64sAreSorted(floats), ShouldBeTrue)
		So(floats, ShouldResemble, []float64{1.2, 3.4, 5.6, 7.8, 9.0})
		So(sort.SearchFloat64s(floats, 6.6), ShouldEqual, 3)
	})

	Convey("test sort slice", t, func() {
		ints := []int{2, 4, 1, 7, 9, 3, 5, 8, 6}
		So(sort.SliceIsSorted(ints, func(i, j int) bool {
			return ints[i] < ints[j]
		}), ShouldBeFalse)
		sort.Slice(ints, func(i, j int) bool {
			return ints[i] < ints[j]
		})
		sort.SliceStable(ints, func(i, j int) bool {
			return ints[i] < ints[j]
		})
		So(sort.SliceIsSorted(ints, func(i, j int) bool {
			return ints[i] < ints[j]
		}), ShouldBeTrue)
		So(ints, ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		So(sort.Search(len(ints), func(i int) bool {
			return ints[i] >= 4
		}), ShouldEqual, 3)
	})

	Convey("test sort interface", t, func() {
		{
			ints := DescIntSlice([]int{2, 4, 1, 7, 9, 3, 5, 8, 6})
			sort.Sort(ints)
			So(ints, ShouldResemble, DescIntSlice([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
		}
		{
			ints := DescIntSlice([]int{2, 4, 1, 7, 9, 3, 5, 8, 6})
			sort.Stable(ints)
			So(ints, ShouldResemble, DescIntSlice([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
		}
	})
}
