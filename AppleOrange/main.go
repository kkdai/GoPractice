package main

import "fmt"

func main() {
	var totalDis [2]int64
	fmt.Scanf("%v %v", &totalDis[0], &totalDis[1])
	if totalDis[0] > totalDis[1] {
		temp := totalDis[1]
		totalDis[1] = totalDis[0]
		totalDis[0] = temp
	}

	var treeDis [2]int64
	fmt.Scanf("%v %v", &treeDis[0], &treeDis[1])
	if treeDis[0] > treeDis[1] {
		temp := treeDis[1]
		treeDis[1] = treeDis[0]
		treeDis[0] = temp
	}

	var totalApple, totalOrange int64
	fmt.Scanf("%v %v", &totalApple, &totalOrange)

	var appleAddress []int64
	for i := 0; int64(i) < totalApple; i++ {
		var appleDis int64
		fmt.Scanf("%v", &appleDis)
		appleAddress = append(appleAddress, treeDis[0]+appleDis)
	}

	var orangeAddress []int64
	for i := 0; int64(i) < totalOrange; i++ {
		var orangeDis int64
		fmt.Scanf("%v", &orangeDis)
		orangeAddress = append(orangeAddress, treeDis[1]+orangeDis)
	}

	appleOut := 0
	for i := 0; i < len(appleAddress); i++ {
		if appleAddress[i] >= totalDis[0] && appleAddress[i] <= totalDis[1] {
			appleOut++
		}
	}
	fmt.Println(appleOut)

	orangeOut := 0
	for i := 0; i < len(orangeAddress); i++ {
		if orangeAddress[i] >= totalDis[0] && orangeAddress[i] <= totalDis[1] {
			orangeOut++
		}
	}
	fmt.Println(orangeOut)
}
