package main

import (
	"fmt"
)

type NothingAtAll []struct{}

func main() {
	//var name string
	//if name == nil {   //ERROR
	//	fmt.Println("the string is nil")
	//}

	var name1 func()

	if name1 == nil {
		fmt.Println("the name1 is nil")
	}

	var nillable chan int

	if nillable == nil {
		fmt.Println("the nillable is nil")
	}

	var e NothingAtAll

	if e == nil {
		fmt.Println("NothingAtAll is  Nil")
	}

	fmt.Println("## nil value:")
	var x []struct{}
	fmt.Printf("type(val): %#v\n", x)
	check := x == nil
	fmt.Println("does value == nil", check)
	foo(x)

	fmt.Println("\n## nil type:")
	fmt.Printf("type(val): %#v\n", nil)
	fmt.Println("does value == nil", true)
	//var y chan int
	foo(nil)

	/**
		在函数返回值中 如有判断nil的情况下 在函数里判断后在返回
		func test() error {
			var p *data = nil
			return p
		}
		error是一个接口类型，test方法中返回的指针p虽然数据是nil，
		但是由于它被返回成包装的error类型，也即它是有类型的。
		所以它的底层结构应该是(*data, nil)，很明显它是非nil的
		func (OsFs) Open(name string) (File, error) {
		    f, e := os.Open(name)
		    if f == nil {
			return nil, e
		    }
		    return f, e
		}
	 */

}

func foo(in interface{}) {
	if in != nil {
		fmt.Println("Not Nil")
	} else {
		fmt.Println("Nil")
	}
}
