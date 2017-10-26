package main

import (
	"bytes"
)

var (
	chans     [4]chan int      //创建生产者
	files     [4]*bytes.Buffer //创建接受者
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
		}(chans[i],i)
	}

}
