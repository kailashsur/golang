package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Welcome to type converion")
	fmt.Println("Please rate our pizza 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseFloat(input, 64);

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Rating is ", numRating
	}

	fmt.Printf("Type of input is %T", input)
}
