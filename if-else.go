package main 

import (
	"fmt"
	"math"
)

func pow (x,n,lim float64) float64 {
	if v := math.Pow(x,n); v < lim {
		return v
	}else{
		fmt.Printf("%g >= %g\n",v,lim)
	}
	return lim

}

func main() {
	fmt.Println(
	pow(3,2,10),
	pow(4,2,12))
	

	fmt.Println("counting")
	
	for i:=0; i<10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
