package main

import (
	"sync"
	"fmt"
)

/**
	sync WaitGroup
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
