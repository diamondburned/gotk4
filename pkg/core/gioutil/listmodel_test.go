package gioutil

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

type weirdType struct {
	weird string
	ptr   *string
}

var weirdTypeList = NewListModelType[weirdType]()

func TestListModel(t *testing.T) {
	expect := []weirdType{
		{weird: "weird1", ptr: ptrTo("weird1")},
		{weird: "weird2", ptr: ptrTo("weird2")},
		{weird: "weird3", ptr: ptrTo("weird3")},
	}

	list := weirdTypeList.New()
	list.ConnectItemsChanged(func(position, removed, added uint) {
		t.Logf(
			"ItemsChanged: position=%d, removed=%d, added=%d",
			position, removed, added)
	})

	fuckShitUp()

	list.Splice(0, 0, expect[0], expect[1])
	list.Append(expect[2])

	fuckShitUp()

	listItems := drainIterator(list.AllItems())
	assertEq(t, "ListModel length mismatch", list.NItems(), 3)
	assertEq(t, "ListModel's items don't match expected list", listItems, expect)

	list.Remove(0)

	listItems = drainIterator(list.AllItems())
	assertEq(t, "ListModel length mismatch", list.NItems(), 2)
	assertEq(t, "ListModel's items don't match expected list", listItems, expect[1:])

	list.Splice(0, 2)

	listItems = drainIterator(list.AllItems())
	assertEq(t, "ListModel length mismatch", list.NItems(), 0)
	assertEq(t, "ListModel's items don't match expected list", listItems, []weirdType(nil))
}

func BenchmarkListModel(b *testing.B) {
	b.Run("append and remove", func(b *testing.B) {
		list := NewListModel[*weirdType]()

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			list.Append((*weirdType)(nil))
			list.Remove(0)
		}
	})

	for _, n := range []int{10, 100, 1000} {
		b.Run(fmt.Sprintf("splice and remove %d", n), func(b *testing.B) {
			list := NewListModel[*weirdType]()
			items := make([]*weirdType, n)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				list.Splice(0, 0, items...)
				list.Splice(0, n)
			}
		})

		b.Run(fmt.Sprintf("iterate over %d", n), func(b *testing.B) {
			list := NewListModel[*weirdType]()
			items := make([]*weirdType, n)
			list.Splice(0, 0, items...)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				all := list.AllItems()
				all(func(*weirdType) bool { return true })
			}
		})
	}
}

func assertEq[T any](t *testing.T, msg string, got, expect T) {
	t.Helper()
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf(msg+"\n"+
			"got:  %#v\n"+
			"want: %#v", got, expect)
	}
}

func drainIterator[T any](iter func(yield func(T) bool)) []T {
	var items []T
	iter(func(item T) bool {
		items = append(items, item)
		return true
	})
	return items
}

func fuckShitUp() {
	for i := 0; i < 10; i++ {
		runtime.GC()
	}
}

func ptrTo[T any](v T) *T {
	return &v
}
