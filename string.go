package main

import "fmt"

func main() {
	var emptystring string

	fmt.Println("empty is:", emptystring)
	fmt.Println("emptystring len", len(emptystring))

	str := "i love you"
	fmt.Println("str :", str)
	fmt.Println("str len:", len(str))
	fmt.Println("str 1:", str[0])

	str = `i
		love
			you`
	fmt.Println("str:", str)

	str = "i like ggggg"
	fmt.Println("str:", str)
	fmt.Println("len str:", len(str))

	rune_str := []rune(str)
	byte_str := []byte(str)

	fmt.Println("len rune str ", len(rune_str))
	fmt.Println("len byte str ", len(byte_str))

	rune_str[7] = 'g'
	rune_str[8] = 'x'
	rune_str[9] = 'j'

	fmt.Println("after ", string(rune_str))

	v_char := 'g'

	fmt.Printf("v_char type %T:", v_char)
}
