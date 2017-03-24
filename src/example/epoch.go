package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()

	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))

	ticker := time.NewTicker(time.Millisecond * 1000)
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at ", t)
		}
	}()

	time.Sleep(time.Millisecond * 5600)
	ticker.Stop()
	fmt.Println("ticker stopped")

	fmt.Println(rand.Intn(100), ',')
	fmt.Println(rand.Intn(100))
	fmt.Println()
	fmt.Println(rand.Float64())

	fmt.Println((rand.Float64()*5)+5, ',')

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ',')

}
