package main

import (
	"fmt"
	"unicode/utf8"
)

type 小学生 struct {
}

func main() {
	fmt.Println("i am  a function !")

	const str = "one world 世界大同"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
}
