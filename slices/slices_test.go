package slices

import (
	"fmt"
)

// ExampleSameUnderlying an example test case
func ExampleSameUnderlying() {
	//create two slices
	first := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	second := first[2:6]
	third := []int{10, 11, 12, 13, 15}
	//check if the first slice is a subset of the second slice
	//and have the same underlying array
	underlying := SameUnderlying(first, second)
	fmt.Println("slices have same underlying array:", underlying)
	underlyingTwo := SameUnderlying(first, third)
	fmt.Println("slices have same underlying array:", underlyingTwo)
	// Output:
	// slices have same underlying array: true
	// slices have same underlying array: false

}
