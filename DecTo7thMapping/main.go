package main

import (
	"fmt"
	"strconv"
)

func main() {
	letters := []string{"0", "a", "b", "c", "d", "e", "f"}

	num := inputNum()
	slice := goSevenDigit(num)
	for i := len(slice) - 1; i >= 0; i-- {
		fmt.Printf("%s", letters[slice[i]])
	}
}

func goSevenDigit(in int64) []int {
	var ret []int
	for in > 0 {
		s := in % 7
		ret = append(ret, int(s))
		in = in / 7
	}
	return ret
}

func inputNum() int64 {
	str := input()
	s, err := strconv.ParseInt(str, 10, 64)
	for err != nil {
		fmt.Println("input is not number")
		str = input()
		s, err = strconv.ParseInt(str, 10, 64)
	}

	return s
}

func input() string {
	var input string
	fmt.Scanln(&input)
	return input
}
