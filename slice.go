package main

import (
	"fmt"
)

var (
	g_str  string = "dada"
	g_int  int    = 32
	g_bool bool   = true
)

const const_str string = "golang"

func main() {
	v_slice := []int{0, 0, 0}
	fmt.Println("v_slice: ", v_slice)

	v_slice[1] = 10
	fmt.Println("v_slice: ", v_slice)

	v_slice = make([]int, 3, 3)
	v_slice[2] = 10

	fmt.Println("v_slice: ", v_slice)

	vp_slice := new([]int)
	fmt.Println("vp_slice: ", vp_slice)
	fmt.Println("vp_slice_len: ", len(*vp_slice))
	fmt.Println("vp_slice_capacity: ", cap(*vp_slice))

	var a string = "golang"
	fmt.Printf(a)

	var d = true
	fmt.Println(d)

	var v_int, int_value = "v_int is:", 2
	fmt.Println(v_int, int_value)

	var b, c int = 2, 3
	fmt.Println(b, c)

	var e int
	fmt.Println(e)

	f := "short item"
	fmt.Println(f)

	v_slice1 := make([]string, 3)
	fmt.Println("empty slice: ", v_slice1)
	v_slice = nil
	fmt.Println("nil slice: ", v_slice1)

	fmt.Println(g_int, g_bool, g_str)

	fmt.Println(const_str)
}
