package main

import "fmt"

type PeopleGetter interface {
	GetName() string
	GetAge() string
}

type Employee struct {
	name   string
	age    int
	salary int
	gender string
}

func (self *Employee) GetName() string {
	return self.name
}

func (self *Employee) GetAge() int {
	return self.age
}

func main() {

	var varEmptyInterface interface{}
	fmt.Printf("varEmptyInterface is of type %T\n", varEmptyInterface)

	varEmptyInterface = 100
	fmt.Printf("varEmptyInterface is of type %T\n", varEmptyInterface)

	varEmptyInterface = "golang"
	fmt.Printf("varEmptyInterface is of type %T\n", varEmptyInterface)

	varEmployee := Employee{
		name:   "jack",
		age:    50,
		salary: 1000,
		gender: "nale",
	}

	fmt.Println("varEmployee is :", varEmployee)
	fmt.Println("varEmployee is :", varEmployee.GetName())
	fmt.Println("varEmployee is :", varEmployee.GetAge())
}
