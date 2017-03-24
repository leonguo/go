package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Result struct {
	Person []Person `xml:"person"`
}

type Person struct {
	Name      string    `xml:"name,attr"`
	Age       int       `xml:"age,attr"`
	Career    string    `xml:"career"`
	Interests Interests `xml:"interests"`
}

type Interests struct {
	Interest []string `xml:"interest"`
}

func main() {
	content, err := ioutil.ReadFile("D://gowork/go/study.xml")
	if err != nil {
		fmt.Printf("err %v", err)
		return
	}
	var result Result
	err = xml.Unmarshal(content, &result)
	if err != nil {
		fmt.Printf("err %v", err)
		return
	}

	for _, data := range result.Person {
		fmt.Printf("name %v", data.Name)

		fmt.Printf("age %v", data.Age)

		fmt.Printf("interests %v %v", data.Interests.Interest[0], data.Interests.Interest[1])

	}
}
