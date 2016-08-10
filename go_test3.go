package main

import "fmt"

func main() {
	c := make(chan int)

	defer close(c)

	go func() { c <- 3 + 4 }()

	i := <-c
	fmt.Println(i)
}
