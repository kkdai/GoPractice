package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Top() int {
	if h.Len() > 0 {
		return (*h)[0]
	}
	return 0
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Top() int {
	if h.Len() > 0 {
		return (*h)[0]
	}
	return 0
}

func add(min *MinHeap, max *MaxHeap, in int) {
	if min.Len() == 0 {
		heap.Push(min, in)
	} else {
		if min.Len() > max.Len() {
			if min.Top() > in {
				ret := heap.Pop(min)
				heap.Push(max, ret)
				heap.Push(min, in)
			} else {
				heap.Push(max, in)
			}
		} else {
			if max.Top() >= in {
				heap.Push(min, in)
			} else {
				ret := heap.Pop(max)
				heap.Push(min, ret)
				heap.Push(max, in)
			}
		}
	}
}

func main() {
	var num int
	fmt.Scanf("%v", &num)

	minList := &MinHeap{}
	maxList := &MaxHeap{}
	heap.Init(minList)
	heap.Init(maxList)

	var outputSlice []float64

	for i := 0; i < num; i++ {
		var temp int
		fmt.Scanf("%v", &temp)
		add(minList, maxList, temp)

		var tempOut float64
		n := minList.Len() + maxList.Len()
		if n%2 == 0 {
			m1 := float64(minList.Top())
			m2 := float64(maxList.Top())
			tempOut = (m1 + m2) / float64(2)
		} else {
			tempOut = float64(minList.Top())
		}

		outputSlice = append(outputSlice, tempOut)
	}

	for i := 0; i < num; i++ {
		fmt.Printf("%1.1f\n", outputSlice[i])
	}
}
