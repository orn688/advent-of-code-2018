package util

// A StringHeap implements heap.Interface and holds strings.
type StringHeap []HeapElement

// A HeapElement assigns a priority to a string for storage and comparison in a
// StringHeap.
type HeapElement struct {
	Value    string
	Priority int
}

func (h StringHeap) Len() int {
	return len(h)
}

func (h StringHeap) Less(i, j int) bool {
	return h[i].Priority < h[j].Priority
}

func (h StringHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push a string onto the end of the heap.
func (h *StringHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's
	// length, not just its contents.
	*h = append(*h, x.(HeapElement))
}

// Pop a string from the top of the heap.
func (h *StringHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
