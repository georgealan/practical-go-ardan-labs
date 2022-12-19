package main

import (
	"fmt"
	"time"
)

func main() {
	// The Go runtime doesn't wait for Goroutines, with that only main will be printed.
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		/* BUG: All goroutines use the same "i" for the loop
		go func() {
			fmt.Println(i)
		}()
		*/

		/* FIX 1: Use a parameter
		go func(n int) {
			fmt.Println(n)
		}(i)
		*/

		// FIX 2: Use a loop body variable
		i := i // "i" shadows "i" from the for loop
		go func() {
			fmt.Println(i) // Use i from the line 27
		}()
	}

	/* Using channel for communication
	Channel semantics:
		- send & receiver will block until opposite operation (*)
		- receive from a closed channel will return the zero value without blocking.
		- send to a closed channel will panic.
		- closing a closed channel will panic too, and doesn't a way to discover if a channel are already closed.
		- send/receive to a nil channel will block forever
	*/
	ch := make(chan string)
	go func() {
		ch <- "hi" // send
	}()
	msg := <-ch // receive
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("got:", msg)
	}

	// Channel is closed, but it won't cause blocking
	msg = <-ch // channel is closed
	fmt.Printf("Closed: %#v\n", msg)

	msg, ok := <-ch // channel is closed
	fmt.Printf("Closed: %#v (ok=%v)\n", msg, ok)

	//ch <- "Hi" // channel is closed and will panic

	// Adding a time lapsing, th egourotine will appear
	time.Sleep(10 * time.Millisecond)
	shadowExample()
}

func shadowExample() {
	n := 7
	{
		n := 2 // From here to } this is inner "n"
		// n = 2 // This is only attribution, don't shadow
		fmt.Println("inner", n)
	}
	fmt.Println("outer", n)
}
