package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(from, ":", i)
	}
}
func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		time.Sleep(time.Second * 1)
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "ping2"
		messages <- "ping3"
	}()

	msg := <-messages
	msg2 := <-messages
	msg3 := <-messages
	fmt.Println(msg, msg2, msg3)

}
