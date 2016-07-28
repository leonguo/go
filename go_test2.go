package main

import "fmt"

//常量的定义
const i = 1

const (
	s  = "hello world"
	s1 = "test"
)

const j1, j2, j3 = 2, 'a', 'b'

//变量

var i1 = 120

var i2, i3 = 22, "dd"

var i4, i5 int

func main() {

	s3 := "hello world"
	fmt.Println(s3)
}
