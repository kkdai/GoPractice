package main

import (
	"fmt"
)

func main() {
	var numSlice, rotation int
	fmt.Scanf("%v %v", &numSlice, &rotation)

	var targetSlice []int
	for i := 0; i < numSlice; i++ {
		var input int
		fmt.Scanf("%v", &input)
		targetSlice = append(targetSlice, input)
	}

	for i := 0; i < numSlice; i++ {
		index := (i + rotation) % numSlice
		fmt.Printf("%d ", targetSlice[index])
	}
}
