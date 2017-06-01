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

type safeCounter1 struct {
	number int
	sync.Mutex
}

func (sc *safeCounter1) Increment(w *sync.WaitGroup) {
	sc.Lock()
	sc.number++
	sc.Unlock()
	defer w.Done()

}
func (sc *safeCounter1) Decrement(w *sync.WaitGroup) {
	sc.Lock()
	sc.number--
	sc.Unlock()
	defer w.Done()
}
func (sc *safeCounter1) getNumber() int {
	sc.Lock()
	number := sc.number
	sc.Unlock()
	return number
}
func main() {
	sc := new(safeCounter1)
	w := sync.WaitGroup{}
	for i := 0; i < 15; i++ {
		w.Add(1)
		go sc.Increment(&w)
		w.Add(1)
		go sc.Decrement(&w)
	}
	w.Wait()
	fmt.Println(sc.getNumber())
}
