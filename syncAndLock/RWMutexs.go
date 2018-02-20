package syncAndLock

import (
	"sync"
	"time"
	"math/rand"
	"fmt"
)

type SynchronizedMapCounter struct {
	m map[int]int
	sync.RWMutex
}

func runWriters(mc *SynchronizedMapCounter, n int) {
	for i := 0; i < n; i++ {
		mc.Lock() // **** Lock Write
		mc.m[i] = i * 10
		mc.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func runReaders(mc *SynchronizedMapCounter, n int) {
	for {
		mc.RLock()	// **** Lock Read
		v := mc.m[rand.Intn(n)] // [0] till [n-1]
		mc.RUnlock()
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}
}


func Start() {
	mapCounter := SynchronizedMapCounter{
		m:make(map[int]int),
	}
	go runWriters(&mapCounter, 10)
	go runReaders(&mapCounter, 10)
	go runReaders(&mapCounter, 10)
	time.Sleep(15 * time.Second)
}
