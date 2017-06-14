package main

import (
	"strings"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

/**
	把中文变成拼音字母 提取首个字母
 */
import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"strings"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getFirstWord(str string) string {
	if len(str) < 1 {
		return ""
	}
	a := pinyin.NewArgs()
	data := pinyin.Pinyin(str, a)
	firstWord := data[0][0]
	firstWord = string([]rune(firstWord)[:1])
	return strings.ToUpper(firstWord)
}

type Filmarea struct {
	Id_  bson.ObjectId `bson:"_id"`
	Area string `json:"area" bson:"area"`
}

func main() {
	// 连接mongo
	session, err := mgo.Dial("192.168.11.200:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("filminfo").C("filmarea")
	result := []Filmarea{}
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range result {
		fmt.Println(v.Id_)
		fmt.Println(v.Area)
		word := getFirstWord(v.Area)
		c.Update(bson.M{"_id": v.Id_}, bson.M{"$set": bson.M{"word": word}})
	}
}
