package routines

import (
	"time"
	"fmt"
)

func Main() {
	go waitAndSay("World")
	fmt.Println("Hello")
	time.Sleep(3 * time.Second)

	// another simplified way (function inside a function)
	go func(s string) {
		time.Sleep(2 * time.Second)
		fmt.Println(s)
	}("World")
	time.Sleep(3 * time.Second)

	word := "Uno"
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(word) // It will print 'Dos'
	}()
	word = "Dos"
	time.Sleep(3 * time.Second)
}

func waitAndSay(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}