package main

import (
	"fmt"
)

func main() {
	var stair int
	fmt.Scanf("%v", &stair)
	if stair <= 0 {
		fmt.Println("cannot input small than 0")
	}

	for i := 1; i <= stair; i++ {
		for j := stair - i; j > 0; j-- {
			fmt.Printf(" ")
		}
		for k := 0; k < i; k++ {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
}
