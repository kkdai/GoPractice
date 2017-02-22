package main

import (
	"fmt"
)

func main() {
	var retAry []int
	totalSize := getSize()

	for i := 0; i < totalSize; i++ {
		for j := 0; j < totalSize; j++ {
			retAry = append(retAry, getNum())
		}
	}

	//Compute Result
	var first, second int
	for k := 0; k < totalSize; k++ {
		first = first + retAry[k+k*totalSize]
	}

	for k := 0; k < totalSize; k++ {
		second = second + retAry[totalSize-1-k+k*totalSize]
	}
	var ret int
	ret = first - second
	if ret < 0 {
		ret = second - first
	}
	fmt.Println(ret)
}

func getSize() int {
	var ret int
	for ret <= 0 {
		ret = getNum()
	}
	return ret
}

func getNum() int {
	var ret int
	fmt.Scanf("%v", &ret)
	return ret
}
