package main

import (
	"fmt"
	"time"
)

func main() {

	// 简单使用 channel 的例子
	varChan1 := make(chan string)
	go func() {
		varChan1 <- "I am a string."
	}()
	fmt.Println("Got msg: ", <-varChan1)

	// 创建 channel 类型对象时设置了 buffer 值
	varChan2 := make(chan string, 3)
	go func() {
		varChan2 <- "input 1"
		varChan2 <- "input 2"
		fmt.Println("I will be before all inputs")
		varChan2 <- "input 3"
	}()
	fmt.Println("Got msg: ", <-varChan2)
	fmt.Println("Got msg: ", <-varChan2)
	fmt.Println("Got msg: ", <-varChan2)

	// channel 中的同步机制
	varChan3 := make(chan bool, 1)
	go func(varChan chan bool) {
		fmt.Println("begin to execute")
		// do something
		time.Sleep(time.Second * 2)
		fmt.Println("end")
		varChan <- true
	}(varChan3)
	fmt.Println("Got msg: ", <-varChan3)

	// 单向 channel 类型对象的使用
	inChan := make(chan string, 1)
	outChan := make(chan string, 1)
	inChan <- "Input str"
	anonymousFunc := func(varInChan chan<- string, varOutChan <-chan string) {
		varInChan <- <-varOutChan
	}
	anonymousFunc(outChan, inChan)
	fmt.Println("Got msg: ", <-outChan)

	// https://gobyexample.com/select
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// https://gobyexample.com/timeouts
	c3 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c3 <- "result 1"
	}()
	select {
	case res := <-c3:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
	c4 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c4 <- "result 2"
	}()
	select {
	case res := <-c4:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

	// https://gobyexample.com/non-blocking-channel-operations
	messages := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	// https://gobyexample.com/closing-channels
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done

	// https://gobyexample.com/range-over-channels
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}
