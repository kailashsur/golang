package main

import "fmt"

type Empoyee struct {
	name string
	age int
	salary int
	id int
}

func main(){
	var p1 Empoyee
	var p2 Empoyee

	// Intialization of the struct
	p1.name = "Kailash"
	p1.age = 20
	p1.salary = 50000
	p1.id = 1


	p2.name = "Sili"
	p2.age = 19
	p2.id = 2
	p2.salary = 40000

	fmt.Println("Name = ", p1.name)
	fmt.Println("Age = ", p1.age)
	fmt.Println("salary = ", p1.salary)
	fmt.Println("ID = ", p1.id)

	fmt.Println("Name = ", p2.name)
	fmt.Println("Age = ", p2.age)
	fmt.Println("salary = ", p2.salary)
	fmt.Println("ID = ", p2.id)
}
