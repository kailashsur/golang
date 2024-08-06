package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to input"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the reading for pizza : ")

	// comma ok || error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for reading, ", input)
	fmt.Printf("Type of reating is %T \n", input)

	// Type conversion in go ********************************

}
