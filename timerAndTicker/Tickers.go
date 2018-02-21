package timerAndTicker

import (
	"time"
	"fmt"
)

func Main() {
	ticker := time.NewTicker(2 * time.Second)

	fmt.Println("Starting Ticker, every 2 seconds...")
	done := tickerCount(ticker)
	time.Sleep(9 * time.Second)
	fmt.Println("Stopping Ticker...")
	ticker.Stop()
	done <- true
	time.Sleep(5 * time.Second)
}

func tickerCount(ticker *time.Ticker) chan bool{
	done := make(chan bool)
	go func() {
		i := 0
	loop:
		for {
			select {
			case t := <-ticker.C:
				i++
				fmt.Println("Count", i , "at", t)
			case <- done : // this channel was added to exiting correctly for this 'for loop'
				fmt.Println("done signal")
				break loop

			}
		}
	}()
	fmt.Println("Exiting the tick counter")
	return done
}