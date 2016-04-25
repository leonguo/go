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

func main() {
	var data = []byte(`
	{"Name":"jemy","Age":26}
	`)

	var s singer
	/**
	docoder := json.NewDecoder(bytes.NewReader(data))
	docoder.Decode(&s)
	docoder.Decode(&t)
	fmt.Println(s)
	fmt.Println(t)
	*/
	err := json.Unmarshal(data, &s)
	fmt.Println(err, s)

}
