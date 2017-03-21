package main

import (
	"fmt"
)

type stack []byte

func (s stack) Push(v byte) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, byte) {
	l := len(s)
	if l == 0 {
		return s, 0
	}
	return s[:l-1], s[l-1]
}

func main() {

	var num int
	fmt.Scanf("%v", &num)

	var inList []string
	for i := 0; i < num; i++ {
		var temp string
		fmt.Scanf("%v", &temp)
		inList = append(inList, temp)
	}

	for _, v := range inList {
		bracket := stack{}

		isBalanced := true
		for _, cha := range []byte(v) {
			if cha == 91 || cha == 123 || cha == 40 {
				bracket = bracket.Push(cha)
			} else if cha == 93 || cha == 125 || cha == 41 {
				var out byte
				bracket, out = bracket.Pop()
				// fmt.Println("pop:", string(out))
				if (cha == 93 && out != 91) || (cha == 125 && out != 123) || (cha == 41 && out != 40) {
					isBalanced = false
					break
				}
			}
		}

		if isBalanced {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")

		}

	}
}
