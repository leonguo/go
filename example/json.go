package main

import (
	"encoding/json"
	"fmt"
)

type singer struct {
	Name string
	Age  int
}

type teacher struct {
	Student string
	Count   int
}

type T struct {
	A  bool
	B  int    "myb"
	B1 int    "2"
	C  string "myc,omitempty"
	D  string `bson:",omitempty" json:"jsonkey"`
	E  int64  ",minsize"
	F  int64  "myf,omitempty,minsize"
	G  string `json:"jsonkey,omitempty"`
}

func main() {
	var data = []byte(`
	{"Name":"jemy","Age":26}
	`)
	var s singer
	err := json.Unmarshal(data, &s)
	fmt.Println(err, s)

	t := &T{G: "0"}
	bt, err := json.Marshal(t)
	fmt.Println(err, string(bt))
}
