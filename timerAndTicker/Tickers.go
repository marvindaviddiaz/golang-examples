package timerAndTicker

import (
	"time"
	"fmt"
)

func Main() {
	ticker := time.NewTicker(2 * time.Second)
	fmt.Println("Starting Ticker, every 2 seconds...")
	go tickerCount(ticker)
	time.Sleep(9 * time.Second)
	ticker.Stop()
	fmt.Println("Stopping Ticker...")
	time.Sleep(5 * time.Second)
	fmt.Println("Exiting...")
}

func tickerCount(ticker *time.Ticker) {
	i := 0
	for t := range ticker.C {
		i++
		fmt.Println("Count", i , "at", t)
	}
}