If a channel is closed no more data can be sent through it.
If we want to know if a channel is closed we can use the second return value of a receive operation.
```go
package main

func main()  {
	ch := make(chan int)
	_, ok := <-ch
	if !ok {
		// channel is closed
	}
}

```
Thi is only if we want to read, but what if we want to write?
```go
package main

func main()  {
    ch := make(chan int)
    ch <- 1 // panic: send on closed channel
}
```
As we can see we can't write to a closed channel, but we can read from it.
But if I need to write to a closed channel?
For this case I wrote a function that wil detect the closed trying to write and catch the panic.
We can use it like this:

```go

package main

import (
	"fmt"
	"github.com/doITmagic/go-snippets/channels/detectClosedChannel"
)

func main() {
	ch := make(chan interface{})
	go func() {
		fmt.Println(<-ch)
	}()
	fmt.Println("channel is closed: ", detectClosedChannel.IsClosed(ch))
	close(ch)
	fmt.Println("channel is closed: ", detectClosedChannel.IsClosed(ch))
}

```
the result will be:
```go
{}
channel is closed:  false
channel is closed:  true

```