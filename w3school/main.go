package main

import "fmt"

func main() {
	// *********** Go variables
	/*
			In Go, there are different types of variables, for example:

		int- stores integers (whole numbers), such as 123 or -123
		float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
		string - stores text, such as "Hello World". String values are surrounded by double quotes
		bool- stores values with two states: true or false

	*/

	var x int = 5
	var y = "string file"
	z := true

	fmt.Println("type is", x)
	fmt.Println("type is", y)
	fmt.Println("type is", z)

	// *********** Go Constants
	const pi = 3.14159
	fmt.Println("value of pi is", pi)

	const (
		pi1 = 3.14159
		a   = 43
		b   = "string foke"
	)
	fmt.Println("value of pi is", pi1)
	fmt.Println("value of a is", a)
	fmt.Println("value of b is", b)

	// *********** Go Arrays **********
	// An array is a collection of items stored at contiguous memory locations. The idea is to store multiple items of the same type together.
	//

	// define aray with size or length of the array
	var arr1 = [5]int{1, 2, 3, 4, 5}
	arr2 := [3]int{10, 21, 23}

	fmt.Println("array 1 is", arr1)
	fmt.Println("array 2 is", arr2)
	fmt.Println("array 1 is", arr1[0])
	fmt.Println("array 2 is", arr2[1])

	// array defining without size
	// defining dynamic array

	var arr3 = [...]string{"a", "b", "c", "d", "e"}
	fmt.Println("array 3 is", arr3)
	fmt.Println("array 3 is", arr3[2])

	arr4 := [...]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println("array 4 is", arr4)
	fmt.Println("array 4 is", arr4[2])

	// array initialization
	arr5 := [5]int{}              //not initialized
	arr6 := [5]int{1, 2}          //partially initialized
	arr7 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println("array 5 is", arr5)
	fmt.Println("array 6 is", arr6)
	fmt.Println("array 7 is", arr7)

	arr8 := [5]int{1: 10, 2: 20} //initialize index 1 and 2
	fmt.Println("array 8 is", arr8)

	// *********** Array Functions **********
	fmt.Println("Length of array :", len(arr8))

}
