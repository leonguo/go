package main

import (
	"fmt"
	"time"
	"unsafe"
)

var start = make(chan bool)
var start1 = make(chan struct{})

func fun1() {
	start <- true
	fmt.Println("This is Worker fun1 ")
}
func fun2() {
	v := <-start
	fmt.Println("This is Worker fun2 ", v)
}

func main() {
	var s struct{}
	fmt.Println(unsafe.Sizeof(s)) // prints 0
	/**
	channel 信号量传递
	*/
	go fun1()
	go fun2()
	time.Sleep(time.Millisecond * 10)
}
