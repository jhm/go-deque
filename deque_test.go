package deque

import (
	"reflect"
	"testing"
)

func TestAddFirst(t *testing.T) {
	elements := []int{3, 2, 1}
	reversed := reverse(elements)
	q := NewWithSize[int](1)
	for i, e := range elements {
		q.AddFirst(e)
		want := reversed[len(reversed)-i-1:]
		if got := q.AsSlice(); !reflect.DeepEqual(got, want) {
			t.Errorf("AsSlice\n got: %v\nwant: %v", got, want)
		}
	}
}

func TestPeekFirst(t *testing.T) {
	q := New[int]()
	if got, found := q.PeekFirst(); got != 0 || found {
		t.Errorf("PeekFirst on empty deque\n got: %d, %v\nwant: 0, false", got, found)
	}
	q.AddLast(1)
	q.AddLast(2)
	q.AddFirst(3)
	if got, found := q.PeekFirst(); got != 3 || !found {
		t.Errorf("PeekFirst\n got: %d, %v\nwant: 3, true", got, found)
	}
}

func TestPeekLast(t *testing.T) {
	q := New[int]()
	if got, found := q.PeekLast(); got != 0 || found {
		t.Errorf("PeekLast on empty deque\n got: %d, %v\nwant: 0, false", got, found)
	}
	q.AddFirst(1)
	q.AddFirst(2)
	q.AddLast(3)
	if got, found := q.PeekLast(); got != 3 || !found {
		t.Errorf("PeekLast\n got: %d, %v\nwant: 3, true", got, found)
	}
}

func TestRemoveFirst(t *testing.T) {
	q := New[int]()
	if got, found := q.RemoveFirst(); got != 0 || found {
		t.Errorf("RemoveFirst on empty deque\n got: %d, %v\nwant: 0, false\n", got, found)
	}
	elements := []int{1, 2, 3}
	for _, e := range elements {
		q.AddLast(e)
	}
	for i, want := range elements {
		if got, found := q.RemoveFirst(); got != want || !found {
			t.Errorf("RemoveFirst\n got: %d, %v\nwant: %d, true", got, found, want)
		}
		if got, want := q.Len(), len(elements)-i-1; got != want {
			t.Errorf("Len\n got: %d\nwant: %d", got, want)
		}
		if got, want := q.AsSlice(), elements[i+1:]; !reflect.DeepEqual(got, want) {
			t.Errorf("AsSlice\n got: %v\nwant: %v", got, want)
		}
	}
}

func TestAddLast(t *testing.T) {
	elements := []int{1, 2, 3}
	q := New[int]()
	for i, e := range elements {
		q.AddLast(e)
		want := elements[:i+1]
		if got := q.AsSlice(); !reflect.DeepEqual(got, want) {
			t.Errorf("AsSlice\n got: %v\nwant: %v", got, want)
		}
	}
}

func TestRemoveLast(t *testing.T) {
	q := New[int]()
	if got, found := q.RemoveLast(); got != 0 || found {
		t.Errorf("RemoveLast on empty deque\n got: %d, %v\nwant: 0, false\n", got, found)
	}
	elements := []int{1, 2, 3}
	reversed := reverse(elements)
	for _, e := range elements {
		q.AddFirst(e)
	}
	for i, want := range elements {
		if got, found := q.RemoveLast(); got != want || !found {
			t.Errorf("RemoveLast\n got: %d, %v\nwant: %d, true", got, found, want)
		}
		if got, want := q.Len(), len(elements)-i-1; got != want {
			t.Errorf("Len\n got: %d\nwant: %d", got, want)
		}
		if got, want := q.AsSlice(), reversed[:len(elements)-i-1]; !reflect.DeepEqual(got, want) {
			t.Errorf("AsSlice\n got: %v\nwant: %v", got, want)
		}
	}
}

func TestAsSlice(t *testing.T) {
	q := NewWithSize[int](1)
	want := []int{}
	if got := q.AsSlice(); !reflect.DeepEqual(got, want) {
		t.Errorf("AsSlice on empty deque\n got: %v\nwant: %v\n", got, want)
	}
	q.AddLast(1)
	q.AddLast(2)
	q.AddLast(3)
	q.AddLast(4)
	// This will cause head > tail.
	q.AddFirst(5)
	want = []int{5, 1, 2, 3, 4}
	if got := q.AsSlice(); !reflect.DeepEqual(got, want) {
		t.Errorf("AsSlice\n got: %v\nwant: %v\n", got, want)
	}
	q.RemoveLast()
	want = []int{5, 1, 2, 3}
	if got := q.AsSlice(); !reflect.DeepEqual(got, want) {
		t.Errorf("AsSlice\n got: %v\nwant: %v\n", got, want)
	}
	// This will cause head < tail again.
	q.RemoveFirst()
	want = []int{1, 2, 3}
	if got := q.AsSlice(); !reflect.DeepEqual(got, want) {
		t.Errorf("AsSlice\n got: %v\nwant: %v\n", got, want)
	}
}

func reverse[T any](xs []T) []T {
	ys := make([]T, len(xs))
	for i, n := range xs {
		ys[len(xs)-i-1] = n
	}
	return ys
}
