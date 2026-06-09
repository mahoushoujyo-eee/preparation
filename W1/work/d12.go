package work

import (
	"fmt"
	"slices"
)

// slice
var oneSlice []int = make([]int, 0)

//array 
// [2]int != [3]int 
var oneArray [2]int = [2]int{}

func someSliceOp(){
	oneSlice = append(oneSlice, 12)
	oneSlice = append(oneSlice, 13)
	oneSlice = append(oneSlice, 14)

	// [,) there is 0, 1
	oneSlice = oneSlice[:2]
	oneVal := oneSlice[0]
	fmt.Print(oneVal)
}

func SlicesDemo(){
	slices.Reverse(oneSlice)
}


// value is ok
func myReverse(s []int){
	n := len(s)
	for i := 0; i < n/2; i++{
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}

// maybe wrong if out of cap
func myAppend(s []int){
	s = append(s, 5)
}

//not concurrent safe
func someMapOp(){
	myMap := make(map[int]int, 0)

	myMap[1] = 11
	v, ok := myMap[2]
	// make sure one val if exist by "ok"
	if ok {
		fmt.Println(v)
	}

	//iter
	for k, v := range myMap {
		fmt.Println(k, v)
	}
}

func repeatDemo(){
	set := make(map[int]int)

	targetArray := []int{1, 1, 3, 2, 3, 4}
	for i, v := range targetArray {
		set[v] = i
	}

	for k, _ := range set {
		fmt.Println(k)
	}
}