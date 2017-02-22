package main

import (
	"fmt"
	"sort"
)

type num []int32

func (a num) Len() int           { return len(a) }
func (a num) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a num) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	var numSlice num
	for i := 0; i <= 4; i++ {
		numSlice = append(numSlice, 0)
	}

	fmt.Scanf("%v %v %v %v %v", &numSlice[0], &numSlice[1], &numSlice[2], &numSlice[3], &numSlice[4])

	//sort it Go 1.8
	sort.Sort(num(numSlice))

	var mini, max int64
	for i := 0; i < 4; i++ {
		mini = mini + int64(numSlice[i])
	}

	for i := 1; i < 5; i++ {
		max = max + int64(numSlice[i])
	}
	fmt.Println(mini, max)
}
