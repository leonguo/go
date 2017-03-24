package main

import (
	"fmt"
)

func main() {
	v_map := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Println("v_map:", v_map)
	fmt.Println("v_map len:", len(v_map))
	v_map[1] = "one hundred"
	v_map[4] = "one hundred"
	fmt.Println("v_map:", v_map)
	fmt.Println("v_map len:", len(v_map))

	v_mapByMake := make(map[int]string, 10)
	fmt.Println("make map:", v_mapByMake)
	fmt.Println("make map len:", len(v_mapByMake))
	v_mapByMake[0] = "one hundred"
	v_mapByMake[1] = "two hundred"
	fmt.Println("make map", v_mapByMake)

	for key, value := range v_mapByMake {
		fmt.Printf("%d:%s\n", key, value)
	}

	fmt.Println("map 10:", v_mapByMake[10])

	delete(v_mapByMake, 10)
	fmt.Println(v_mapByMake)

	v_mapOfArray := map[int][]int{
		1: {0, 1, 2},
		2: {3, 4, 5},
	}
	fmt.Println(v_mapOfArray)

	fmt.Println("v_map:", v_map)

	for key, _ := range v_map {
		delete(v_map, key)
		if key == 1 {
			v_map[key*5] = "new"
		}
	}
	fmt.Println("after v_map:", v_map)
}
