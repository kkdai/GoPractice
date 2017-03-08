package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inpuStr string
	fmt.Scanln(inpuStr)
	reader := bufio.NewReader(os.Stdin)
	inpuStr, _ = reader.ReadString('\n')
	fmt.Println("Hello, World.")
	fmt.Println(inpuStr)
}
