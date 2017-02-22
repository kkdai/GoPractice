package main

import (
	"fmt"
	"strconv"
)

const (
	//MaxArraySize Need modify this value if your pattern is much learger than 100
	MaxArraySize int = 100
)

func main() {
	firstSize := inputSize()
	firtAry := inputAry(firstSize)
	//fmt.Println(firtAry)

	secondSize := inputSize()
	secondAry := inputAry(secondSize)
	//fmt.Println(secondAry)

	fmt.Println(Strstr(firtAry, secondAry))
}

func Strstr(haystack []string, needle []string) int {
	retSlice := KMP(haystack, needle)
	if len(retSlice) > 0 {
		return retSlice[0]
	}

	return -1
}

func KMP(haystack []string, needle []string) []int {
	next := preKMP(needle)
	i := 0
	j := 0
	m := len(needle)
	n := len(haystack)

	x := needle
	y := haystack
	var ret []int

	//got zero target or want, just return empty result
	if m == 0 || n == 0 {
		return ret
	}

	//want string bigger than target string
	if n < m {
		return ret
	}

	for j < n {
		for i > -1 && x[i] != y[j] {
			i = next[i]
		}
		i++
		j++

		//fmt.Println(i, j)
		if i >= m {
			ret = append(ret, j-i)
			//fmt.Println("find:", j, i)
			i = next[i]
		}
	}

	return ret
}

func preKMP(x []string) [MaxArraySize]int {
	var i, j int
	length := len(x) - 1
	var kmpNext [MaxArraySize]int
	i = 0
	j = -1
	kmpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = kmpNext[j]
		}

		i++
		j++

		if x[i] == x[j] {
			kmpNext[i] = kmpNext[j]
		} else {
			kmpNext[i] = j
		}
	}
	return kmpNext
}

func inputSize() int {
	firstSize := input()
	iFirstSize, err := strconv.Atoi(firstSize)
	for err != nil {
		fmt.Println("input is not number")
		firstSize = input()
		iFirstSize, err = strconv.Atoi(firstSize)
	}
	return iFirstSize
}

func input() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func inputAry(size int) []string {
	var retAry []string
	for i := 0; i < size; i++ {
		val := input()
		if len(val) > 0 {
			retAry = append(retAry, val)
		} else {
			i--
		}

	}

	return retAry
}
