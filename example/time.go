package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(time.Second)

	d := <-time1.C
	fmt.Println("time 1 expired", d)

	time2 := time.NewTimer(time.Second)

	go func() {
		<-time2.C
		fmt.Println("time 2 expired")
	}()

	time.Sleep(10 * time.Second)

	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("time 2 stopped")
	}
	fmt.Println("end")

}
