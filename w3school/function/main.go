package main

import (
	"fmt"
)

// normal function defination and declaration
func message() {
	fmt.Println("hello, how are you")
}

// function with multi parameter
func parameter_func(name string, age int) {
	fmt.Println("Name is ", name)
	fmt.Println("Age is ", age)
}

// functin return type and return statement
func add_two_num(num1 int, num2 int) int { // in this statement, there is int return type, that means, it should return only int type of data

	var resutl int = num1 + num2
	fmt.Println("Result is returned ", resutl)
	return resutl
}

// named return value syntax
func named_return(x int, y int) (res int) {
	res = x + y

	fmt.Println("res returned ", res)
	return
}

// multiple return value
func multi_return (x int, y int, name string)(res int, name_res string){
	res = x+y
	name_res = name
	return
}


func main() {
	// calling the function inside the main function

	message()
	message()

	parameter_func("Kailash", 20)

	var total int = add_two_num(20, 30)
	fmt.Println("Total result is ", total)

	// named return value
	res_val := named_return(10, 20)
	fmt.Println("res_val = ", res_val)

	// multiple return value
	res1, res2 := multi_return(3, 5, "kailash");
	fmt.Println("First val = ",res1)
	fmt.Println("Second val = ", res2)
}
