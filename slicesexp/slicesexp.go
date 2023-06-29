package slicesexp

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func Example() {
	s := []int{1, 2, 3}
	s2 := []int{1, 2, 3., 5}

	// Check if slice contain 2
	contain := slices.Contains(s, 2)
	fmt.Printf("slice %v contain %d : %v \n", s, 2, contain)

	// Check if slice contains multiple of value
	contains := slices.Contains(s, 3)
	fmt.Printf("slice %v contain %d and %d : %v \n", s, 2, 3, contains)

	//us equals to check if slice contain all values
	contain = slices.Equal(s, s2)
	fmt.Printf("slice %v equals to %v : %v \n", s, s2, contain)

	//use grow to increase the size of the slice
	s = slices.Grow(s, 2)
	fmt.Printf("slice %v grow by 2 : %v \n", s, s2)

	//use insert to insert a value at a given index
	s = slices.Insert(s, 2, 4)
	fmt.Printf("slice %v insert 4 at index 2 : %v \n", s, s2)

	// use to replace a value at a given index
	s = slices.Replace(s, 2, 4)
	fmt.Printf("slice %v replace value at index 2 by 4 : %v \n", s, s2)

	// use to compact a slice
	s = slices.Compact(s)
	fmt.Printf("slice %v compact : %v \n", s, s2)

	// use to determine if is sorted
	contain = slices.IsSorted(s)
	fmt.Printf("slice %v is sorted : %v \n", s, contain)

	//use to clip a slice , remove unused capacity
	s = slices.Clip(s)
	fmt.Printf("slice %v clip : %v \n", s, s2)

}
