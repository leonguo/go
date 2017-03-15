package main

import "fmt"

type User struct {
	UserId   int
	UserName string
}

func SimplePanicRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic info is ", err)
		}
	}()
	panic("SimplePanicRecover function ed")
}

func GetUser() (user User) {
	user = User{}
	user.UserId = 1
	user.UserName = "gg"
	return
}

func main() {
	SimplePanicRecover()
	test := fmt.Sprintf("test ddd %d", 2)
	fmt.Printf(test)

	user := GetUser()
	fmt.Println(user.UserId)
	fmt.Println(user)
}
