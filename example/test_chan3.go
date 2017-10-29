package main

import (
	"time"
	"fmt"
)

/**
	定时器
 */

var ch = make(chan int)

func main() {

	mytimer := time.NewTicker(time.Millisecond * 1000)
	go func() {
		ch <- 1
	}()
	for {
		select {
		case <-ch:
			go func() {
				fmt.Println(" task ")
			}()
		case <-mytimer.C:
			go func() {
				fmt.Println("my time ticker")
			}()
		}
	}
	fmt.Println("Hello go")

	time.Sleep(time.Millisecond * 2000)
}
