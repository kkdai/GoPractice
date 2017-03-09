package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank"

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var firstI uint64
	var secodD float64
	var thirdS string

	// Read and save an integer, double, and String to your variables.
	// var pos int64 = 0
	// scanLines := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 	advance, token, err = bufio.ScanLines(data, atEOF)
	// 	pos = pos + int64(advance)
	// 	return
	// }
	// scanner.Split(scanLines)

	scanner.Scan()
	target := scanner.Text()
	firstI, _ = strconv.ParseUint(target, 10, 64)

	scanner.Scan()
	target = scanner.Text()
	var err error
	secodD, err = strconv.ParseFloat(string(target), 64)
	if err != nil {
		fmt.Println(err)
	}
	// ss := "4.0"
	// fmt.Println([]byte(target), []byte(ss), "-----", secodD)

	scanner.Scan()
	thirdS = scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + firstI)

	// Print the sum of the double variables on a new line.
	fmt.Printf("%0.1f\n", d+secodD)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + thirdS)

}
