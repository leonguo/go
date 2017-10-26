package main

import (
	"bytes"
	"strconv"
	"fmt"
)

const STEP, NUM = 4, 40

var (
	chans     [STEP]chan int      //创建生产者
	files     [STEP]*bytes.Buffer //创建接受者 一个缓存byte类型的缓冲器存放着都是byte
	last, cur = 0, 0
)

func main() {
	for i := range chans {
		chans[i] = make(chan int)
		files[i] = new(bytes.Buffer)
		go func(out chan<- int, prime int) {
			for {
				out <- prime
			}
		}(chans[i], i)
	}
	for j := 0; j < NUM; j++ {
		k := <-chans[j%STEP] + 1
		files[cur].WriteString(strconv.Itoa(k))
		// 当前游标如果到达最大值 则停留一次，然后继续前进
		if k != STEP || cur == last {
			last, cur = cur, (cur+1)%4
		}
	}
	fmt.Print(files)
}
