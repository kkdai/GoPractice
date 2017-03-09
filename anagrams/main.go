package main

import (
	"fmt"
)

func main() {
	var alphabetListA, alphabetListB [26]int
	var strA, strB string

	strA = getStr()
	strB = getStr()

	alphabetListA = countAlphabet(strA)
	alphabetListB = countAlphabet(strB)

	comapreResult(alphabetListA, alphabetListB)

}

func getStr() string {
	var str string
	fmt.Scanf("%s", &str)
	return str
}

func countAlphabet(str string) [26]int {
	var ret [26]int
	for i := 0; i < len(str); i++ {
		valInt := (int)(str[i] - 97)
		// fmt.Println(valInt)
		ret[valInt]++
	}
	return ret
}

func comapreResult(a [26]int, b [26]int) {
	var result int
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			diff := a[i] - b[i]
			if diff < 0 {
				diff = 0 - diff
			}
			result = result + diff

		}
	}
	fmt.Println(result)
}
