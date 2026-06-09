package base

import "fmt"

var oneVal = 1

func init() {
	fmt.Printf("p1 init %d\n", oneVal)
}

func P1() {
	println("p1")
}