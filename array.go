package main

import (
	"fmt"
)

func printArrayValue(arr [1]int) {
	fmt.Println("arr[0] is ", &arr[0])
}

func printAarrayref(arr *[1]int) {
	fmt.Println("arr[0] address", &arr[0])
}

func main() {
	v_IntArray := [...]int{}
	fmt.Printf("%T\n", v_IntArray)
	fmt.Println("len arr is", len(v_IntArray))
	fmt.Println("cap arr is", cap(v_IntArray))

	v_IntArrayOf5 := [5]int{}
	fmt.Println("len(v_IntArrayOf5)", len(v_IntArrayOf5))
	fmt.Println("v_IntArrayOf5", v_IntArrayOf5)

	v_IntArrayOf5[2] = 3
	fmt.Println("v_IntArrayOf5", v_IntArrayOf5)

	v_IntArrayOf5 = [5]int{1, 2}
	fmt.Println("len arr:", len(v_IntArrayOf5))
	fmt.Println("v_IntArrayOf5", v_IntArrayOf5)

	v_IntArrayOf5 = [5]int{3, 1, 3: 4, 10}
	fmt.Println("len(v_IntArrayOf5: ", len(v_IntArrayOf5))
	fmt.Println("v_IntArrayOf5: ", v_IntArrayOf5)

	v_IntArrayOf1 := [1]int{}
	fmt.Println("address arr[0]", &v_IntArrayOf1[0])
	fmt.Println(v_IntArrayOf1)

	printArrayValue(v_IntArrayOf1)
	printAarrayref(&v_IntArrayOf1)

	v_IntArrayof23 := [2][3]int{}
	fmt.Println(v_IntArrayof23)

	v_IntArray1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	v_IntSlice := v_IntArray1[1:4:9]
	fmt.Println("v_IntArray[1:4:9]", v_IntSlice)
	fmt.Println("v_IntArray[1:4:9] len ", len(v_IntSlice))
	fmt.Println("v_IntArray[1:4:9] len ", cap(v_IntSlice))

	v_IntSlice = v_IntArray1[0:]
	v_IntSlice[0] = 100
	fmt.Println(v_IntSlice)
	fmt.Println("len: ", len(v_IntSlice))
	fmt.Println("len: ", cap(v_IntSlice))

	v_IntSlice = v_IntArray1[:1]
	fmt.Println(v_IntSlice)
	fmt.Println("len: ", len(v_IntSlice))
	fmt.Println("len: ", cap(v_IntSlice))

	v_IntSlice = v_IntArray1[:4:4]
	fmt.Println(v_IntSlice)
	fmt.Println("len: ", len(v_IntSlice))
	fmt.Println("len: ", cap(v_IntSlice))

	v_IntSlice = make([]int, 5, 10)
	v_IntSliceAnother := v_IntSlice[3:]
	v_IntSliceAnother[0] = 100
	fmt.Println(v_IntSliceAnother)
	v_IntSliceAnother = append(v_IntSlice, 100)
	fmt.Println(v_IntSliceAnother)
	fmt.Println("make([]int, 5, 10), v_IntSlice is: ", v_IntSlice)
	fmt.Println("make([]int, 5, 10), len(v_IntSlice) is: ", len(v_IntSlice))
	fmt.Println("make([]int, 5, 10), cap(v_IntSlice) is: ", cap(v_IntSlice))
}
