package main

import (
	"fmt"
)

func main() {
	var numSlice int
	fmt.Scanf("%v", &numSlice)

	var targetSlice []int
	for i := 0; i < numSlice; i++ {
		var input int
		fmt.Scanf("%v", &input)
		targetSlice = append(targetSlice, input)
	}

	for i := 0; i < numSlice; i++ {
		fmt.Printf("%d ", targetSlice[numSlice-i-1])
	}
}
