package work

import (
	"errors"
	"fmt"
)

type oneS struct {
	val1 int
	val2 int
}

func newDemo() {
	s1 := new(oneS)
	s1.val1 = 1

	s2 := &oneS{val1: 1, val2: 2}
	s2.val2 = 3
}

func ErrorDemo() error {
	err1 := errors.New("error content")
	err2 := fmt.Errorf("error with val: %d", 0)

	//what is %w
	allErrs := fmt.Errorf("err1: %w, err2: %w", err1, err2)
	errors.Is(allErrs, err1)

	return allErrs
}