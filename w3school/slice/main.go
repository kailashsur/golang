package main

import "fmt"

/*
	Go 'Slice' is similar to the array but It is more flaxible and powerful than array

	Like arrays, slices are also used to store multiple values of the same type in a single variable.

However, unlike arrays, the length of a slice can grow and shrink as you see fit.

In Go, there are several ways to create a slice:

Using the []datatype{values} format
Create a slice from an array
Using the make() function
*/

func main() {
	fmt.Println("Go Slice")

	// common way to declare a slice
	// var slice1 = []int{}
	// slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}

	// there are 2 most popular function to find the length of the slice
	//len() function - returns the length of the slice (the number of elements in the slice)
	// cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

	fmt.Println("Length of slice3: ", len(slice3))
	fmt.Println("Capacity of slice3: ", cap(slice3))

	// ****create a slice from an array
	//
	// Syntax
	// var myarray = [length]datatype{values} // An array
	// myslice := myarray[start:end] // A slice made from the array

	arr1 := [5]int{1, 2, 3, 4, 5}
	slice4 := arr1[1:4]
	fmt.Println("Slice4: ", slice4)
	fmt.Println("Length of slice4: ", len(slice4))
	fmt.Println("Capacity of slice4: ", cap(slice4))

	// ***** create slice using make() function
	// Syntax
	// make([]datatype, length, capacity)
	// The make() function creates a slice of the specified length and capacity.
	// The length is the number of elements the slice contains, and the capacity is the number of elements the slice can hold.
	// The capacity argument is optional. If you don't specify the capacity, the capacity will be equal to the length.
	// The make() function initializes the slice with zero values.
	// The make() function is commonly used to create a slice that can grow or shrink in size.

	slice5 := make([]int, 5, 10)
	// slice6 := make([]int , 3); // ommited capacity

	fmt.Println("Slice using make () = ", slice5) // it initialize all place at 0
	fmt.Println("Length of slice5 : ", len(slice5))
	fmt.Println("Capacity of sllie 5 : ", cap(slice5))

	// ***************************Slice Modification **********************
	/*
		severeal modificaton in go
		1. access in slice
		2. change of slice
		3. append of in slice
		4. copy of slice
	*/

	// 1. Access value of slice
	arr7 := []int{1, 2, 3, 4, 5}
	slice7 := arr7[2:3]
	fmt.Println("Slice is : ", slice7)
	fmt.Println("Value : ", slice7[0])

	// 2. change element of slice
	fmt.Println("Before change element : ", slice7[0])
	slice7[0] = 5
	fmt.Println("After change element of slice : ", slice7[0])

	// 3. Append element to the slice
	// 		you can add element to the end of the slice by append

	var slice8 = []int{1, 2, 3, 4, 5}

	fmt.Println("Beffore append, full slice = ", slice8)

	newSlice := append(slice8, 6, 7, 8)
	fmt.Println("After append, slice = ", newSlice)

	// 4. Append all elements to one slice to another slice
	// syntax :
	// total_slice = append(slice1, slice2...)

	slice9 := []int{1, 2, 3, 4, 5}
	slice10 := []int{6, 7, 8, 9, 0}

	total_slice := append(slice9, slice10...)
	fmt.Println("First slice = ", slice9)
	fmt.Println("First slice = ", slice10)
	fmt.Println("total append slice = ", total_slice)

	/*
		5. copy() in slice

		The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.

		syntax :
		copy(dist, src)

		The copy() function takes in two slices dest and src, and copies data from src to dest. It returns the number of elements copied.

	*/

	src_slice := slice9
	dist_slice := make([]int, 10)

	fmt.Println("Before dist slice = ", dist_slice)

	copy(dist_slice, src_slice)
	fmt.Println("After dist slice = ", dist_slice)

}
