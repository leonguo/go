package main

import "fmt"

type EmputyStruct struct {
}

type Job struct {
	title string "JD"
	rank  int
}

type JobWithoutTag struct {
	title string
	rank  int
}

func (Job) Help() {
	fmt.Println("get detail help")
}

func (job Job) What() {
	fmt.Printf("detail: title:%s\n rank:%d\n", job.title, job.rank)
}

func (jobP *Job) SetRank(newRank int) {
	jobP.rank = newRank
	fmt.Println("new value is ", jobP.rank)
}

type Employee struct {
	_                   int
	age                 int
	firstName, lastName string
	pack                struct {
		salary int
		stock  int
	}
	*Job
}

// 定义 `struct` 类型，这个类型唯一的属性为空接口类型，而任何类型都实现了空接口
// 所以 OnlyInterface 可以实例化属性为任何类型的 `struct` 对象
type OnlyInterface struct {
	f interface{}
}

func main() {
	varEmputyStruct := EmputyStruct{}
	fmt.Println("emputy struct", varEmputyStruct)

	var varJob0 Job
	fmt.Println("varJob0", varJob0)

	varJob1 := Job{
		"ceo",
		99,
	}
	fmt.Println("varJob1", varJob1)

	varJob2 := Job{
		rank:  99,
		title: "ceo",
	}

	fmt.Println("varJob2", varJob2)

	varJob3 := Job{
		title: "coo",
	}

	fmt.Println("varJob3", varJob3)

	// Employee  `struct` 类型对象的默认值
	varEmployee0 := Employee{}
	fmt.Println("varEmployee = ", varEmployee0)

	// 初始化 Employee `struct` 类型对象，匿名属性需要在初始化其他属性后使用对象名显式赋值
	varEmployee1 := Employee{
		age:       50,
		firstName: "Jack",
		lastName:  "Ma",
	}
	varEmployee1.pack.salary = 100000000
	varEmployee1.pack.stock = 1000000
	varJob4 := Job{
		"Founder",
		100,
	}
	varEmployee1.Job = &varJob4
	fmt.Println("varEmployee = ", varEmployee1)
	fmt.Println("varEmployee.Job = ", *(varEmployee1.Job))

	// 注意 Employee 的 Help 方法 `override` 了其属性 Job 的 Help 方法
	varEmployee1.Help()
	// 调用 Job 的 Help 方法需要显式调用
	varEmployee1.Job.Help()
	// 使用 Employee 对象可以直接调用其属性 Job 的方法 What
	varEmployee1.What()
	// Employee 对象的属性 Job 的类型是 `*Job`，所以可以修改 Job 的内容
	// 如果其类型是 Job 本身，由于获取到的 `struct` 对象是其原始值的拷贝，修改不会生效
	varEmployee1.Job.SetRank(88)
	fmt.Println("After job rank change, varEmployee.Job = ", *(varEmployee1.Job))

	// `struct` 类型对象是值类型，可以进行比较运算
	varEmployee2 := Employee{
		age:       50,
		firstName: "Jack",
		lastName:  "Ma",
	}
	varEmployee2.pack.salary = 100000000
	varEmployee2.pack.stock = 1000000
	varEmployee2.Job = &varJob4
	if fmt.Println("Cmpare 2 struct object."); varEmployee2 == varEmployee1 {
		fmt.Println("varEmployee2 equals varEmployee1.")
	}

	// Tag 是 `struct` 类型的一部分，下面 varJob5 无法赋值给 varJob4
	// varJob5 := JobWithoutTag{
	// 	"Founder",
	// 	100,
	// }
	// varJob4 = varJob5

	// 属性类型为空接口的 `struct` 类型的使用
	varOnlyInterface := OnlyInterface{
		f: 100,
	}
	fmt.Println("OnlyInterface with int: ", varOnlyInterface)
	varOnlyInterface = OnlyInterface{
		f: "I like Golang.",
	}
	fmt.Println("OnlyInterface with string: ", varOnlyInterface)

}
