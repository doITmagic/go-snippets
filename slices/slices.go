package slices

import (
	"fmt"
)

// SameUnderlying check if the first slice and the second slice have the same underlying array
func SameUnderlying[T any](first, second []T) bool {
	if len(first) == 0 && len(second) == 0 {
		return false
	}

	//verify if the address of the elements of first slice are equal to the address of one of the element of the second slice
	for i := 0; i < len(first); i++ {
		firstAddress := fmt.Sprintf("%p", &first[i])
		for j := 0; j < len(second); j++ {
			if firstAddress == fmt.Sprintf("%p", &second[j]) {
				//address of the element of the first slice is equSWal to the address of one of the element of the second slice
				return true
			}
		}
	}

	return false
}
