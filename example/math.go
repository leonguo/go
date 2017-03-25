package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

const (
	Big = 1 << 1 // 位运算 左移动1位
)

func main() {
	fmt.Println(add(2, 56))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
