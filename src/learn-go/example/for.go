package main

import (
	"fmt"
)

func main() {
	v_str := "test str"
	fmt.Printf(v_str)

	for i := 0; i < 7; i++ {
		fmt.Printf("v_str[%d]: %q\t", i, v_str[i])
	}
	fmt.Println()
	for i, length := 0, len(v_str); i < length; i++ {
		fmt.Printf("v_str[%d]: %q\t", i, v_str[i])
	}
	fmt.Println()

	for i, elem := range v_str {
		elem = elem - 32
		fmt.Printf("v_str[%d]:%q\t", i, elem)
	}
	fmt.Printf(v_str)

	j := 0
	for j < 4 {
		fmt.Printf("v_str[%d]:%q\t", j, v_str[j])
		break
	}
	fmt.Println()
	count := 0
	for {
		if count > 3 {
			break
		}
		fmt.Printf("count %d", count)
		count++
	}

	v_IntArray := [...]int{1, 2, 3, 4, 5}

	for i, elem := range v_IntArray {
		elem *= 100
		fmt.Println(i, elem)
	}
	fmt.Println(v_IntArray)

	for i, elem := range &v_IntArray {
		elem *= 100
		fmt.Println(i, elem)
	}
	fmt.Println(v_IntArray)

	v_PointIntArray := [...]*int{&v_IntArray[0], &v_IntArray[1]}
	for i, elem := range v_PointIntArray {
		*elem *= 100
		fmt.Println(i, *elem)
	}
	fmt.Println(v_PointIntArray)

	for _, elem := range v_PointIntArray {
		fmt.Printf("%d\t", *elem)
		fmt.Println(*elem)
	}
	fmt.Println()

	v_int1 := 9
	if v_int1 < 8 {
		fmt.Printf("true")
	}

	switch v_int1 {
	case v_IntArray[0]:
		fmt.Printf("test")
	default:
		fmt.Printf("default")
	}
}
