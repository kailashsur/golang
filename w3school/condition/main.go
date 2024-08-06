package main

import "fmt"

func main() {

	x := 20
	y := 30

	if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("y is greater than x")

		if x == 20 {
			fmt.Println("x is ", x)
		}
	}

	// switch case conditions

	var day int = 4

	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("mon")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println(" day is 4")
	case 5:
		fmt.Println("day is 5")
	default:
		fmt.Println("Invalid day")
	}

}
