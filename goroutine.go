package main

import (
	"fmt"
	"runtime"
	"sync"
)

func FuncA(wg *(sync.WaitGroup)) {
	defer wg.Done()
	fmt.Println("this is funcA")
}

type DummyStruct struct {
}

func (ds *DummyStruct) FuncB(wg *(sync.WaitGroup)) {
	defer wg.Done()
	fmt.Println("This is DummyStruct.FuncB().")
}

// 当以 goroutine 方式运行时主动中断自身
func FuncWithTermination(wg *(sync.WaitGroup)) {
	defer func() {
		fmt.Println("defer in FuncWithTermination")
		wg.Done()
	}()
	runtime.Goexit()
	fmt.Println("This is FuncWithTermination.")
}

// 当以 goroutine 方式运行时暂停自身之后系统线程空闲后再恢复运行
func FuncWithGosched(wg *(sync.WaitGroup)) {
	defer wg.Done()
	for index := 0; index < 5; index++ {
		fmt.Println("Count ", index)
		if index == 2 {
			runtime.Gosched()
		}
	}
}

func main() {
	// goroutine 的个数
	grCount := 5

	// 设置可以并行运行的核心数
	runtime.GOMAXPROCS(1)

	// 等待 goroutine 执行结果
	wg := new(sync.WaitGroup)
	wg.Add(grCount)
	fmt.Println("Starting...")
	go FuncWithGosched(wg)
	go FuncA(wg)
	ds := DummyStruct{}
	go ds.FuncB(wg)
	go func(wg *(sync.WaitGroup)) {
		defer wg.Done()
		fmt.Println("This is inner anonymous func.")
	}(wg)
	go FuncWithTermination(wg)
	wg.Wait()
}
