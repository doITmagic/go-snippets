package detectClosedChannel

import (
	"fmt"
)

func ExampleIsClosed() {
	ch := make(chan interface{})
	go func() {
		fmt.Println(<-ch)
	}()
	fmt.Println("channel is closed: ", IsClosed(ch))
	close(ch)
	fmt.Println("channel is closed: ", IsClosed(ch))
	// Output:
	//{}
	//channel is closed:  false
	//channel is closed:  true
}
