package main

import (
	"fmt"
)

func main() {
	var numS, numT int
	fmt.Scanf("%v %v", &numS, &numT)

	if numS < numT {
		fmt.Println("No")
		return
	}

	s := InputMap(numS)
	t := InputAry(numT)

	if CompareTwoMap(s, t) == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

//CompareTwoMap :
func CompareTwoMap(s map[string]int, t []string) bool {
	for _, k := range t {
		if s[k] == 0 {
			return false
		} else {
			s[k] = s[k] - 1
		}
	}
	return true
}

///InputMap :
func InputMap(size int) map[string]int {
	m := make(map[string]int)
	for i := 0; i < size; i++ {
		val := Input()
		if len(val) > 0 {
			m[val] = m[val] + 1
		} else {
			i--
		}

	}

	return m
}

//InputAry :
func InputAry(size int) []string {
	var retAry []string
	for i := 0; i < size; i++ {
		val := Input()
		if len(val) > 0 {
			retAry = append(retAry, val)
		} else {
			i--
		}
	}
	return retAry
}

//Input :
func Input() string {
	var input string
	fmt.Scanf("%v", &input)
	return input
}
