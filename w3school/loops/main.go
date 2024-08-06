package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("i is ", i)
	}
	fmt.Println("Reverse loop")
	for i := 5; i > 0; i-- {
		fmt.Println("i is ", i)
	}

	/*
				The "Range" Keyword
			The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.

			The range keyword is used like this:

			Syntax
		for index, value := array|slice|map {
		   // code to be executed for each iteration
		}
	*/

	fruits := [3]string{"apple", "banana", "chery"}

	fmt.Println("Range of fruits array", fruits)
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

}
