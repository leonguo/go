package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

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

var b1 bool

var b2 uint8 //无符号8位整数 2^8 (0-255) byte 的别称

var b3 int8 //无符号8位整数 2^8 (-128-127)

var b4 int32 // rune的别称

func BytesToString(b []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	strHeader := reflect.StringHeader{bytesHeader.Data, bytesHeader.Len}
	return *(*string)(unsafe.Pointer(&strHeader))
}
func StringToBytes(s string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bytesHeader := reflect.SliceHeader{strHeader.Data, strHeader.Len, strHeader.Len}
	return *(*[]byte)(unsafe.Pointer(&bytesHeader))
}

type File interface {
	Close()
}

func main() {
	buffer1 := [10]string{}

	intSet := [6]int{1, 2, 3, 4}

	days := [...]string{"Sat", "Sun"}
	s3 := "hello world"
	fmt.Println(s3)
	b1 = true
	fmt.Println(b1)
	b2 = 12
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)

	fmt.Println(buffer1)
	fmt.Println(intSet)
	fmt.Println(days)

	var s1 []int
	var s2 = make([]int, 10)
	var s4 = make([]int, 10, 20)

	fmt.Printf("len:%d, cap:%d\n", len(s1), cap(s1))
	fmt.Printf("len:%d, cap:%d\n", len(s2), cap(s2))

	fmt.Printf("len:%d, cap:%d\n", len(s3), cap(s4))

	z := "hello world"

	z1 := StringToBytes(z)
	z2 := BytesToString(z1)

	fmt.Println(z1)

	fmt.Printf("%s\n", z2)

	x := 10
	p := &x

	fmt.Println(p)
	//Map是一组无序的键值对的集合。键的类型相同，值的类型也相同。
	m1 := make(map[int]int, 10)
	m1[0] = 1
	m1[1] = 2
	fmt.Println(m1)

}
