package main

import (
	"fmt"
	"time"
)

var start = make(chan bool)

func fun1() {
	start <- true
	fmt.Println("This is Worker fun1 ")
}
func fun2() {
	v := <-start
	fmt.Println("This is Worker fun2 ",v)
}

func main() {
	/**
	 channel 信号量传递
	*/
	go fun1()
	go fun2()
	time.Sleep(time.Millisecond * 10)
}
