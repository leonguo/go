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
		Id bson.ObjectId "_id"
		Seq int
	}
	result := Counter{}
	selecter := bson.M{}
	s.DB("test").C("counters").Find(selecter).One(&result)
	fmt.Println(result.Id)
	fmt.Println(result.Seq)
	// Output:
	// 1
	// 2
	// 3
}