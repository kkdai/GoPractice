package main

import "fmt"

func main() {
	num := GetNum()

	if num%2 == 1 {
		fmt.Println("Weird")
	} else if num >= 6 && num <= 20 {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
}

//GetNum :
func GetNum() int {
	var ret int
	fmt.Scanf("%v", &ret)
	return ret

}
