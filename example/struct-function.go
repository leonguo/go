package main

import "fmt"

type Circle struct {
	radius float64
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
}

func (c *Circle) getArea() float64 {
	fmt.Println(c)
	c.radius = 20.00
	return 3.14 * c.radius * c.radius
}
