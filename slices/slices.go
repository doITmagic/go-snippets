package slices

import (
	"fmt"
)

// IsSubset check if the one of the slice is a subset of another slice
func IsSubset[T any](first, second []T) bool {
	if len(first) == 0 && len(second) == 0 {
		return false
	}

	//loop over the first slice
	for i := 0; i < len(first); i++ {
		//get the address of the element of the first slice
		firstAddress := fmt.Sprintf("%p", &first[i])
		//loop over the second slice
		for j := 0; j < len(second); j++ {
			//get the address of the element of the second slice
			if firstAddress == fmt.Sprintf("%p", &second[j]) {
				//if the address of the element of the first slice is equal to the address of the element of the second slice
				return true
			}
		}
	}

	return false
}

// SameUnderlying check if the first slice and the second slice have the same underlying array
func SameUnderlying[T any](first, second []T) bool {
	//	verify if both slices have a capacity greater than 0
	return cap(first) > 0 && cap(second) > 0 &&
		&first[:cap(first)][cap(first)-1] == &second[:cap(second)][cap(second)-1]
}
