package main

import (
	"sync"
	"fmt"
)

/**
	mutex 互斥锁
	互斥锁是传统并发程序对共享资源进行访问的主要手段。是sync包中的Mutex结构体
	   type Mutex struct{}
	该结构体包括两个方法，可以说是非常简单使用的
	   func (m *Mutex) Lock() {}
	   func (m *Mutex) UnLock() {}
	我们来通过一个简单的例子来说明他的用法
 */
var c chan int

type safeCounter struct {
	number int
	sync.Mutex
}

func (sc *safeCounter) Increment(i int) {
	sc.Lock()
	sc.number++
	sc.Unlock()
	c <- i
}
func (sc *safeCounter) Decrement(i int) {
	sc.Lock()
	sc.number--
	sc.Unlock()
	c <- i
}
func (sc *safeCounter) getNumber() int {
	sc.Lock()
	number := sc.number
	sc.Unlock()
	return number
}
func main() {
	sc := new(safeCounter)
	//w := sync.WaitGroup{}
	c = make(chan int)
	for i := 0; i < 10; i++ {
		go sc.Increment(i)
		fmt.Printf("%v", <-c)
		go sc.Decrement(i)
		fmt.Printf("%v", <-c)
	}
	fmt.Println(sc.getNumber())
}
