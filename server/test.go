package main

import (
	"fmt"
	"godoc/db/mongodb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	c, err := mongodb.Dial("120.24.229.18", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	// session
	s := c.Ref()
	defer c.UnRef(s)
	err = s.DB("test").C("counters").RemoveId("test")
	if err != nil && err != mgo.ErrNotFound {
		fmt.Println(err)
		return
	}

	// auto increment
	err = c.EnsureCounter("test", "counters", "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 3; i++ {
		id, err := c.NextSeq("test", "counters", "test")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id)
	}

	// index
	c.EnsureUniqueIndex("test", "counters", []string{"key1"})

	type Counter struct {
		Id  bson.ObjectId "_id"
		Seq int
	}
	type Users struct {
		Username string "user_name"
		Age      int        "age"
	}
	//result := Counter{}
	//selecter := bson.M{"_id":"test"}
	//s.DB("test").C("counters").Find(selecter).One(&result)
	//fmt.Println(result.Id)
	//fmt.Println(result.Seq)

	users := Users{}
	selecter := bson.M{"user_name":"ggg", "age":20}
	err = s.DB("test").C("users").Insert(selecter)
	fmt.Println(err)

	//selecter = bson.M{"user_name":"ggg"}
	//update := bson.M{"$set":bson.M{"age":21}}
	//err = s.DB("test").C("users").Update(selecter, update)
	//fmt.Println(err)

	selecter = bson.M{"user_name":"ggg"}
	iter := s.DB("test").C("users").Find(selecter).Sort("_id").Iter()
	for iter.Next(&users) {
		fmt.Printf("users name: %v age: %v\n",users.Username,users.Age)
	}

	// Output:
	// 1
	// 2
	// 3
}