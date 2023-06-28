package main

import (
	"fmt"
)

func main() {
	//create slice
	as := []int{1, 2, 3, 4, 5}
	fmt.Println("initial: ", as)
	//send 'as' slice by value
	fmt.Printf("new slice returned %#v \n", newReceiveSliceByValue(as))
	fmt.Println("initial slice after running newReceiveSliceByValue: ", as)
}

// newReceiveSliceByValue receive slice by value but add a return slice value
func newReceiveSliceByValue(s []int) []int {
	//create a new slice with the same length of the original slice
	newS := make([]int, len(s))
	//copy the values of the original slice
	copy(newS, s)
	//change the 5th element of the new slice
	newS[4] = 10
	//return the new slice
	return newS
}
