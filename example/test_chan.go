package main

import (
	"fmt"
)

func worker(start chan bool, index int) {
	<-start
	fmt.Println("This is Worker:", index)
}

func main() {
	/**
	等待一个事件(Event)
	*/
	fmt.Println("begin doing something")
	start := make(chan bool)
	go func() {
		fmt.Println("start doing something")
		close(start)
	}()
	v , ok := <-start
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Println("done")

}
