## What is happening if we use a slice like parameters in a function, sending it by value? It will be changed?
  Let's see an example:
  ```go
  package main

import (
	"fmt"
)

func main() {
	//create slice
	as := []int{1, 2, 3, 4, 5}
	fmt.Println("initial: ", as)
	//send 'as' slice by value
	receiveSliceByValue(as)
	fmt.Println("after: ", as)
}

// receiveSliceByValue receive slice by value
// and change the 5th element
func receiveSliceByValue(s []int) {
	s[4] = 10
}

```
  The result will be:
  ```
 initial:  [1 2 3 4 5]
 after:  [1 2 3 4 10]

  ```

### ok, but why?
  Because a slice is a reference type. <br/>
  To be more clear, we are sending the slice by value, but the copied slice point to the same address of the first element of the copied slice. <br/>
  So, if we change the value of the slice in the function, we are changing the value of the slice in the main function, because both slices point to the same array address. <br/>
  If we want to send a slice by value we need to create a new slice and copy the values of the original slice.
  ```go
  package main

import (
	"fmt"
)

func main() {
	//create slice
	as := []int{1, 2, 3, 4, 5}
	fmt.Println("initial: ", as)
	//send 'as' slice by value
	fmt.Printf("new slice returned %#v \n",newReceiveSliceByValue(as))
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

```

Now the result will be:
```
initial:  [1 2 3 4 5]
new slice returned []int{1, 2, 3, 4, 10}
initial slice after running newReceiveSliceByValue:  [1 2 3 4 5]
```
the initial slice is not changed. <br/>