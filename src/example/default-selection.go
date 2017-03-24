package main

import (
	"fmt"
	"time"
)

func main() {
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

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
