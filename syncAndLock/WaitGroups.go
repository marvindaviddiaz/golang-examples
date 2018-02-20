package syncAndLock

import (
	"sync"
	"time"
	"math/rand"
	"fmt"
)

func Run() {
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {

		//Increment the wait group counter
		wg.Add(1)

		go func (i int) {
			// Decrement the counter when the goroutine completes
			defer wg.Done()

			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Println("Work done for", i)
		}(i)
	}

	wg.Wait() // the program will ended when wait group counter will be '0'

}