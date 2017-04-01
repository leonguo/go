package main

import "fmt"

type Circle struct {
	radius float64
}

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func main() {
	var c Circle
	var d = 10
	di := &d
	fmt.Println("a is ", d)
	fmt.Println("a is ", d)
	fmt.Println("a is ", *di)
	c.radius = 10.00
	fmt.Println("area is :", c.getArea())
	fmt.Println("area is :", c.radius)
	fmt.Println("area is :", &c)

	user := NewUser("guoxingjun", 20)
	fmt.Println(user.Name)
}

func (c *Circle) getArea() float64 {
	fmt.Println(c)
	c.radius = 20.00
	return 3.14 * c.radius * c.radius
}
