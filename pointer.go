package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/**
	 *  *取指针所处的对象
	 *  &取对象的地址
	 */
	var p *int
	fmt.Println("p is ", p)

	pp := &p
	fmt.Printf("type pp %T", pp)
	fmt.Println("value pp ", pp)

	intvar := 1000
	p = &intvar
	fmt.Println("after assignment p is ", p)
	fmt.Println("value p point ", *p)

	var strp *[4]byte
	strp = (*[4]byte)(unsafe.Pointer(p))
	fmt.Println(strp)
	fmt.Println(*strp)

	type user struct {
		name string
	}

	userp := &user{
		"xiaoa",
	}
	fmt.Println("user point to is ", *userp)
	userp.name = "dddd"
	fmt.Println("user point to is ", *userp)

	fmt.Println("test-----------------part2---------")
	
	i,j := 42,299
	ppp := &i
	fmt.Println(*ppp)
	*ppp = 21 
	fmt.Println(i)
	
	ppp = &j
	
	*ppp = *ppp/10
	fmt.Println(j)
	
}
