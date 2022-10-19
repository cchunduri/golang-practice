package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Please give input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Input is ", input)
}
