package main

import (
	"fmt"
)

func main() {
	var mealCost float64
	var tipPercent float64
	var taxPercent float64

	mealCost = GetFloat()
	tipPercent = GetFloat() * 0.01
	taxPercent = GetFloat() * 0.01

	tip := mealCost * tipPercent
	tax := mealCost * taxPercent
	fmt.Printf("The total meal cost is %0.0f dollars.", mealCost+tip+tax)
}

//GetFloat :
func GetFloat() float64 {
	var ret float64
	fmt.Scanf("%v", &ret)
	return ret

}
