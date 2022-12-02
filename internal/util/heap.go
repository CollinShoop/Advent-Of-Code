package util

// An IntHeap is a min-heap of ints.
// https://pkg.go.dev/container/heap#example-package--PriorityQueue

type IntHeap struct {
	data  []int
	isMin bool
}

func NewIntHeap(isMin bool) *IntHeap {
	return &IntHeap{
		data:  make([]int, 10),
		isMin: isMin,
	}
}

func (h IntHeap) Len() int { return len(h.data) }
func (h IntHeap) Less(i, j int) bool {
	if h.isMin {
		return h.data[i] < h.data[j]
	}
	return h.data[i] > h.data[j]
}
func (h IntHeap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.data = append(h.data, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old.data)
	x := old.data[n-1]
	h.data = old.data[0 : n-1]
	return x
}
