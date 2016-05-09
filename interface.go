package main

import "fmt"

type Employee struct {
	name   string
	age    int
	salary int
	gender string
}

func (self *Employee) GetName() string {
	return self.name
}

func (self *Employee) GetAge() int {
	return self.age
}

func (self *Employee) Help() {
	fmt.Println("this is help info")
}

func (self *Employee) GetSalary() int {
	return self.salary
}

func (self *Employee) GetGender() string {
	return self.gender
}

// 匿名接口可以被用作变量或者结构属性类型
type Man struct {
	gender interface {
		GetGender() string
	}
}

// 定义执行回调函数的接口
type Callbacker interface {
	Execute()
}

// 定义函数类型 func() 的新类型 CallbackFunc
type CallbackFunc func()

// 实现 CallbackFunc 的 Execute() 方法
func (self CallbackFunc) Execute() { self() }

// 定义接口类型 PeopleGetter 包含获取基本信息的方法
type PeopleGetter interface {
	GetName() string
	GetAge() int
}

// 定义接口类型 EmployeeGetter 包含获取薪水的方法
// EmployeeGetter 接口中嵌入了 PeopleGetter 接口，前者将获取后者的所有方法
type EmployeeGetter interface {
	PeopleGetter
	GetSalary() int
	Help()
}

func main() {

	var varEmptyInterface interface{}
	fmt.Printf("varEmptyInterface is of type %T\n", varEmptyInterface)

	varEmptyInterface = 100
	fmt.Printf("varEmptyInterface is of type %T\n", varEmptyInterface)

	varEmptyInterface = "golang"
	fmt.Printf("varEmptyInterface is of type %T\n", varEmptyInterface)

	varEmployee := Employee{
		name:   "jack",
		age:    50,
		salary: 1000,
		gender: "nale",
	}

	fmt.Println("varEmployee is :", varEmployee)
	varEmployee.Help()
	fmt.Println("varEmployee is :", varEmployee.GetName())
	fmt.Println("varEmployee is :", varEmployee.GetAge())
	fmt.Println("varEmployee is :", varEmployee.GetSalary())
	fmt.Println("varEmployee is :", varEmployee.GetGender())

	varMan := Man{&Employee{
		name:   "nobody",
		age:    20,
		salary: 1000,
		gender: "unknow",
	}}
	fmt.Println("the gender of nobody is :", varMan.gender.GetGender())

	// 接口类型转换，从超集到子集的转换是可以的
	// 从方法集的子集到超集的转换会导致编译错误
	// 这种情况下 switch 不支持 fallthrough
	var varEmpInter EmployeeGetter = &varEmployee

	switch varEmpInter.(type) {
	case nil:
		fmt.Println("nil")
	case PeopleGetter:
		fmt.Println("peoplegetter")
	default:
		fmt.Println("unknown")
	}

	varCallbacker := CallbackFunc(func() {
		fmt.Println("i am a callback function")
	})

	varCallbacker.Execute()
}
