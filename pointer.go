package main

import "fmt"

func main() {
	var p *int
	fmt.Println("p is ", p)

	pp := &p
	fmt.Printf("type pp %T", pp)
	fmt.Println("value pp ", pp)

	intvar := 1000
	p = &intvar
	fmt.Println("after assignment p is ", p)
	fmt.Println("value p point ", *p)

}
