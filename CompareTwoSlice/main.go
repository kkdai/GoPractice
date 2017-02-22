package main

import "fmt"

func main() {
	a := GetNum()
	b := GetNum()

	var aWin, bWin int
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			aWin++
		} else if a[i] < b[i] {
			bWin++

		}
	}
	fmt.Println(aWin, bWin)

}

func GetNum() (ret [3]int) {

	var a [3]int
	for a[0] <= 0 || a[1] <= 0 || a[2] <= 0 || a[0] > 100 || a[1] > 100 || a[2] > 100 {
		fmt.Scanf("%v %v %v", &a[0], &a[1], &a[2])
	}
	return a
}
