package main

import "fmt"

func fn(m map[int]int) {
	m = make(map[int]int)
	m[213] = 231
	fmt.Printf("fn %v", m)
}

// 值传递
func main() {
	var a int
	var b, c = &a, &a
	println(&b, &c)
	println(b, c)

	var m map[int]int
	fmt.Printf("1 %v", m)
	fn(m)
	fmt.Printf("2 %v", m)
	fmt.Println(m == nil)
}
