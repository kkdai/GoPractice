package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
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

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
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

/*
void add(int i){
    if(lower.empty())
        lower.push(i);
    else{
        if(lower.size() > higher.size()){
            if(lower.top() > i){
                higher.push(lower.top());
                lower.pop();
                lower.push(i);
            }
            else
                higher.push(i);
        }
        else{
            if(higher.top() >= i)
                lower.push(i);
            else{
                lower.push(higher.top());
                higher.pop();
                higher.push(i);
            }
        }
    }
}
double find(){
    int n = lower.size() + higher.size();
    return (n % 2 == 0) ? (higher.top() + lower.top())/2.0 : (double)(lower.top());

}
*/
func add(min *MinHeap, max *MaxHeap, in int) {
	if min.Len() == 0 {
		heap.Push(min, in)
	} else {
		if min.Len() > max.Len() {
			if (*min)[0] > in {
				ret := heap.Pop(min)
				heap.Push(max, ret)
				heap.Push(min, in)
			} else {
				heap.Push(max, in)
			}
		} else {
			if (*max)[0] >= in {
				heap.Push(min, in)
			} else {
				ret := heap.Pop(max)
				heap.Push(min, ret)
				heap.Push(max, in)
			}
		}
	}

	fmt.Println(min, max)
}

func main() {
	var num int
	fmt.Scanf("%v", &num)

	minList := []*MinHeap{}
	maxList := []*MaxHeap{}

	for i := 0; i < num; i++ {
		min := &MinHeap{}
		max := &MaxHeap{}
		heap.Init(min)
		heap.Init(max)

		minList = append(minList, min)
		maxList = append(maxList, max)
	}

	// var outputSlice []float64
	med := &IntHeap{}
	// max := &MaxHeap{}
	heap.Init(med)
	// heap.Init(max)

	for i := 0; i < num; i++ {
		var temp int
		fmt.Scanf("%v", &temp)
		// var ret float64

		heap.Push(med, temp)
		// heap.Push(max, temp)
		if i == 5 {
			for j := 0; j < 5; j++ {
				fmt.Println(heap.Pop(med))
			}
		}
		fmt.Println(med)
		for j := i; j < num; j++ {
			add(minList[j], maxList[j], temp)
		}

		// fmt.Println(minList, maxList)
		// outputSlice = append(outputSlice, ret)
	}

	for i := 0; i < num; i++ {

		n := minList[i].Len() + maxList[i].Len()
		if n%2 == 0 {
			m1 := float64(heap.Pop(minList[i]).(int))
			m2 := float64(heap.Pop(maxList[i]).(int))
			fmt.Printf("%1.1f\n", (m1+m2)/float64(2))
		} else {
			fmt.Printf("%1.1f\n", float64(heap.Pop(minList[i]).(int)))
		}
	}
}
