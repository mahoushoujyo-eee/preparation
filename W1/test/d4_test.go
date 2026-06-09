package test

import (
	"fmt"
	"testing"
	"w1/work"
)

func TestErrorDemo(t *testing.T){
	allErrs := work.ErrorDemo()
	fmt.Println(allErrs)
}