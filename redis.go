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

	values, _ := redis.Values(c.Do("lrange", "redlist", "0", "100"))
	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}
	//或者
	var v1 string
	redis.Scan(values, &v1)

	fmt.Println("rest")
}
