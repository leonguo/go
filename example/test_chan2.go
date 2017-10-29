package main

import (
	"fmt"
	"time"
)
/**
	生产者 消费者
 */
func producer(c chan int, max int) {
	for i := 0; i < max; i++ {
		fmt.Println("i is ",i)
		c <- i
	}
}

func consumer(c chan int) {
	ok := true
	value := 0
	for ok {
		fmt.Println("Wait receive")
		if value, ok = <-c; ok {
			fmt.Println(value)
		}
		if ok == false{
			fmt.Println("*******Break********")
		}
	}
}

func main() {
	c := make(chan int)
	defer close(c)
	go producer(c, 10)
	go consumer(c)
	time.Sleep(time.Second * 5)
	fmt.Println("Done")
}