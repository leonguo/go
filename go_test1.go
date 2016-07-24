package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type 小学生 struct {
}

func main() {
	fmt.Println("i am  a function !")

	const str = "one world 世界大同"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))

	str1 := "on 时间"

	r := []rune(str1)
	fmt.Printf("%#v\n", r)

	str1 = string(r)
	fmt.Printf("%#v\n", str1)

	s := strconv.QuoteRune('☺')
	fmt.Println(s)
	fmt.Println('☺')

}
