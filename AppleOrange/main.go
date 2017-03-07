package main

import "fmt"

func main() {
	var totalDis [2]int64
	totalDis[0], totalDis[1] = getTwoNumber()

	var treeDis [2]int64
	treeDis[0], treeDis[1] = getTwoNumber()

	var totalApple, totalOrange int64
	fmt.Scanf("%v %v", &totalApple, &totalOrange)

	appleAddress := getList(totalApple, treeDis[0])
	orangeAddress := getList(totalOrange, treeDis[1])

	appleOut := calculateFruit(appleAddress, totalDis[0], totalDis[1])
	fmt.Println(appleOut)

	orangeOut := calculateFruit(orangeAddress, totalDis[0], totalDis[1])
	fmt.Println(orangeOut)
}

func calculateFruit(friSlice []int64, base1, base2 int64) int {
	ret := 0
	for i := 0; i < len(friSlice); i++ {
		if friSlice[i] >= base1 && friSlice[i] <= base2 {
			ret++
		}
	}
	return ret
}

func getTwoNumber() (int64, int64) {
	var val1, val2 int64
	fmt.Scanf("%v %v", &val1, &val2)
	if val1 > val2 {
		temp := val2
		val2 = val1
		val1 = temp
	}
	return val1, val2
}

func getList(length int64, base int64) []int64 {
	var retSlice []int64
	for i := 0; int64(i) < length; i++ {
		var dis int64
		fmt.Scanf("%v", &dis)
		retSlice = append(retSlice, base+dis)
	}
	return retSlice
}
