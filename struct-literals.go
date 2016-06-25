package main 

import "fmt"

type Vertex struct {
	X, Y int 
}

var (
	v1 = Vertex{1,2}
	v2 = Vertex{X:1}	
	v3 = Vertex{}
	p = &Vertex{1,2}
)

func main(){
	fmt.Println(v1,p,v2,v3)
	
	a := make([]int, 5)
		
	fmt.Println(a)
	
	var b []int 
	fmt.Println(b)
	b = append(b,0)
	fmt.Println(b)
	b = append(b,1)
	fmt.Println(b)
	b = append(b,2,3,4,5)
	fmt.Println(b)
}
