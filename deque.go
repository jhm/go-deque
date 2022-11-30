// Package deque provides a resizable double-ended queue.
package deque

const DefaultDequeSize = 16

// A double-ended queue.
type deque[T any] struct {
	elements   []T
	head, tail int
}

// New returns a new deque with the default size.
func New[T any]() *deque[T] {
	return NewWithSize[T](DefaultDequeSize)
}

// NewWithSize returns a new deque with the given initial size (rounded up to
// the nearest greater power of 2).
func NewWithSize[T any](initialSize int) *deque[T] {
	size := nextPowerOf2(initialSize)
	return &deque[T]{elements: make([]T, size)}
}

// IsEmpty returns true iff the queue has no items in it.
func (d *deque[T]) IsEmpty() bool {
	return d.head == d.tail
}

// Len returns the number of elements in the queue.
func (d *deque[T]) Len() int {
	return (d.tail - d.head) & (len(d.elements) - 1)
}

// AddFirst adds an element to the front of the queue.
func (d *deque[T]) AddFirst(element T) {
	d.head = (d.head - 1) & (len(d.elements) - 1)
	d.elements[d.head] = element
	if d.head == d.tail {
		d.grow()
	}
}

// AddLast adds an element to the end of the queue.
func (d *deque[T]) AddLast(element T) {
	d.elements[d.tail] = element
	d.tail = (d.tail + 1) & (len(d.elements) - 1)
	if d.head == d.tail {
		d.grow()
	}
}

// PeekFirst returns the element at the front of the queue and a bool indicating
// that an element was found. The zero value for type T and false will be
// returned when there are no elements in the queue.
func (d *deque[T]) PeekFirst() (T, bool) {
	if d.IsEmpty() {
		return zero[T](), false
	}
	return d.elements[d.head], true
}

// PeekLast returns the element at the end of the queue and a bool indicating
// that an element was found. The zero value for type T and false will be
// returned when there are no elements in the queue.
func (d *deque[T]) PeekLast() (T, bool) {
	if d.IsEmpty() {
		return zero[T](), false
	}
	return d.elements[d.tail-1], true
}

// RemoveFirst removes the element at the front of the queue and returns it and
// a bool value indicating that an element was found. The zero value for type T
// and false will be returned when there are no elements in the queue.
func (d *deque[T]) RemoveFirst() (T, bool) {
	if d.IsEmpty() {
		return zero[T](), false
	}
	first := d.elements[d.head]
	d.elements[d.head] = zero[T]()
	d.head = (d.head + 1) & (len(d.elements) - 1)
	return first, true
}

// RemoveLast removes the element at the end of the queue and returns it and a
// bool value indicating that an element was found. The zero value for type T
// and false will be returned when there are no elements in the queue.
func (d *deque[T]) RemoveLast() (T, bool) {
	if d.IsEmpty() {
		return zero[T](), false
	}
	i := (d.tail - 1) & (len(d.elements) - 1)
	last := d.elements[i]
	d.elements[i] = zero[T]()
	d.tail = i
	return last, true
}

// AsSlice returns a slice that includes all the elements in the deque in the
// order they would be returned from repeated calls to RemoveFirst.
func (d *deque[T]) AsSlice() []T {
	return d.copyToSliceWithSize(d.Len())
}

func (d *deque[T]) grow() {
	newSize := len(d.elements) << 1
	xs := d.copyToSliceWithSize(newSize)
	newTail := len(d.elements)
	d.elements = xs
	d.head = 0
	d.tail = newTail
}

func (d *deque[T]) copyToSliceWithSize(size int) []T {
	xs := make([]T, size)
	if d.head < d.tail {
		copy(xs, d.elements[d.head:d.tail])
	} else {
		n := copy(xs, d.elements[d.head:])
		copy(xs[n:], d.elements[:d.tail])
	}
	return xs
}

func zero[T any]() T {
	var z T
	return z
}

func nextPowerOf2(n int) int {
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	return n + 1
}
