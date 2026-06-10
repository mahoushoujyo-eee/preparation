package tool

import "testing"

func TestCat(t *testing.T) {
	err := Cat([]string{"../test.txt"})
	if err != nil {
		t.Fatalf("test fail, %v", err)
	}
}
