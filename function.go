package main

import (
	"fmt"
)

func noReturn() {
	fmt.Println("no return message")
}

func withParamsNoReturn(intParams1, intParams2 int, strParams string) {
	fmt.Println("my params are:", intParams1, intParams2, strParams)

}

func withParamsReturn(intParams int) (int, int) {
	return intParams * 2, intParams
}

func withParasAndReturn(intParams int) (doubledInput int) {
	fmt.Println("inital value of doubledInput:", doubledInput)
	doubledInput = intParams * 2
	return
}

// 定义有参数也有含命名返回值的函数，如果内部声明了和命名返回值同名的局部变量，需要显式返回命名返回值
func WithpNamedReturnAndInnerVar(intPara int) (output int) {
	{
		var output = 100
		output = intPara * 2
		return output
	}
}

func WithVariabicPara(prefix string, paras ...int) {
	temInt := 0
	for _, para := range paras {
		temInt += para
	}
	fmt.Println("%s,the sum of para is: %d\n", prefix, temInt)
}

func withClosure(intPara int) func() {
	returnFunc := func() {
		fmt.Println("Input is :", intPara)
	}
	return returnFunc
}

func main() {
	noReturn()
	withParamsNoReturn(1, 2, "ok")
	input, output := withParamsReturn(10)
	fmt.Println("return :", input, output)
	fmt.Println("the double 10 is:", withParasAndReturn(10))
	fmt.Println("the double 10 is:", WithpNamedReturnAndInnerVar(10))
	WithVariabicPara("test", 1, 2, 3, 4)
	withClosure(100)()
}
