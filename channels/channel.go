package channels

import (
	"fmt"
)

func Main() {

	c := make(chan bool)
	go waitAndSay(c, "World")
	fmt.Println("Hello")

	// We send a signal to c in order to allow waitAndSay to continue
	c <- true

	// we wait to receive another signal on c before we exit
	<- c

	// Using buffered Channels
	bufferedChannel()

	// using range and close
	bufferedChannelUsingRangeAndCloseKeywords()
}

func waitAndSay(c chan bool, s string) {
	if b := <- c; b {
		fmt.Println(s)
	}
	c <- true
}

func bufferedChannel() {

	fmt.Println("Using buffered channel:")

	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func bufferedChannelUsingRangeAndCloseKeywords() {
	fmt.Println("Using buffered Range and Close:")
	c := make(chan string)
	go sayHelloMultipleTimes(c, 5)
	for s := range c {
		fmt.Println(s)
	}

	v, ok := <- c
	fmt.Println("Channel close?", !ok, "value", v)
}

func sayHelloMultipleTimes(c chan string, n int) {
	for i := 0; i < n; i++ {
		c <- "Hello"
	}
	close(c)
}


