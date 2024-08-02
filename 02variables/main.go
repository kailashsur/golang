package main

import "fmt"

var jwtToken = 30000 // this will work, because we are using var keyword to declare the variable
// jwtToken := 30000 // this will not work, because we are not using var keyword to declare the variable,
// we can not use shorthand declaration outside the function body, it will give the error
// only inside the function body we can use the volorious operator

// declaring the constants ********************************
// In Go, constants are declared like variables, but with the const keyword.
// * and first leter should be capital
// * This is public
const LoginToken string = "slkdjfaierrejae" // this is the way to declare the constant in go
// const loginToken string = "sdkfja" // this will not work, because we are not using the capital letter to declare the constant

// Go is a statically typed language, which means that the type of a variable (or a constant, or a function parameter) is known at compile time.
func main() {
	var username string = "Kailash"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLogedIn bool = true
	fmt.Println(isLogedIn)
	fmt.Printf("Variable is of type : %T \n", isLogedIn)

	var smallval uint8 = 255 // 0 to 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type : %T \n", smallval)

	var num int = 25576
	fmt.Println(num)
	fmt.Printf("Variable is of type : %T \n", num)

	var smallFloat float32 = 255.4551111 // float32 is the set of all IEEE-754 32-bit floating-point numbers.
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	var largeFlat float64 = 255.4551111 // float64 is the set of all IEEE-754 64-bit floating-point numbers.
	fmt.Println(largeFlat)
	fmt.Printf("Variable is of type : %T \n", largeFlat)

	// alices
	// default values and some aliases
	var anotherVar int // default value is 0
	fmt.Println(anotherVar)
	fmt.Printf("Variable is type of : %T \n", anotherVar)

	var anotherVar2 string // default value is ""
	fmt.Println(anotherVar2)
	fmt.Printf("Variable is type of : %T \n", anotherVar2)

	var anotherVar3 bool // default value is false
	fmt.Println(anotherVar3)
	fmt.Printf("Variable is type of : %T \n", anotherVar3)

	// implicit type of declaring the variables
	// it is implicitly lexier can understand the type of the variable
	var website = "kailash.com"
	fmt.Println(website)
	fmt.Printf("Variable is type of : %T \n", website)

	// shorthand declaration
	var name, email = "Kailash", "k@gmail.com"
	fmt.Println(name, email)

	// no var style to declare variables
	// this is most common syntax to declare the variables, and most of the developers are useing it in their code base
	numberOfUser := 100000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is type of : %T \n", numberOfUser)

	floatNoUser := 200.0
	fmt.Println(floatNoUser)
	fmt.Printf("Variable is type of : %T \n", floatNoUser)

	// use public var
	fmt.Println(jwtToken)
	fmt.Printf("Variable is type of : %T \n", jwtToken)
	fmt.Println(LoginToken)
	fmt.Printf("Variable is type of : %T \n", LoginToken)
}

// uploading to the github
// all github commands are
// step by steps
// git init
// git add .
// git commit -m "first commit"
// git branch -M main
// git remote add origin
