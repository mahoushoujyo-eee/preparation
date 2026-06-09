package service

import (
	"fmt"
	_ "w1/ano_base"
	"w1/base"
)

func init(){
	fmt.Println("s1 init")
}

func S1() {
	fmt.Println("s1")
	base.P1()
}