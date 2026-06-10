package tool

import "testing"

func TestGrep(t *testing.T) {
	err := Grep([]string{"txt", "../test.txt"})

	if err != nil {
		t.Fatalf("test fail: %v", err)
	}
}
