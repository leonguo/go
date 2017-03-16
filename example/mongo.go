package main

import (
	"fmt"
)

func main() {
	fmt.Println("mongotest")
	/**
	   查询所有
	   var users []Users.Find(nil).All(&users)
	    =($eq)

	         c.Find(bson.M{"name": "Jimmy Kuu"}).All(&users)

	    !=($ne)

	         c.Find(bson.M{"name": bson.M{"$ne": "Jimmy Kuu"}}).All(&users)

	    >($gt)

		c.Find(bson.M{"age": bson.M{"$gt": 32}}).All(&users)

	    <($lt)

		c.Find(bson.M{"age": bson.M{"$lt": 32}}).All(&users)

	    >=($gte)

		c.Find(bson.M{"age": bson.M{"$gte": 33}}).All(&users)

            <=($lte)

		c.Find(bson.M{"age": bson.M{"$lte": 31}}).All(&users)

	    in($in)

		c.Find(bson.M{"name": bson.M{"$in": []string{"Jimmy Kuu", "Tracy Yu"}}}).All(&users)
	   and($and)

	        c.Find(bson.M{"name": "Jimmy Kuu", "age": 33}).All(&users)

	   or($or)

		c.Find(bson.M{"$or": []bson.M{bson.M{"name": "Jimmy Kuu"}, bson.M{"age": 31}}}).All(&users)
	 */
}
