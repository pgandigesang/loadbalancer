package main

type MinHeap []*Server

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i].CurrentLoad < (*h)[j].CurrentLoad
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Server))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
