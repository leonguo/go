package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := c.Do("AUTH", "leon"); err != nil {
		c.Close()
		return
	}
	v, err := c.Do("SET", "name", "red")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	fmt.Println(v)

	fmt.Println("rest")
}
