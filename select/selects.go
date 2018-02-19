package selects

import (
	"time"
	"math/rand"
	"fmt"
)

var crewMapping = map[string]int {
	"James": 5,
	"Kevin": 10,
	"Rahul": 9,
}

func findCrew(name, server string, c chan int) {
	// Simulate searching
	time.Sleep(time.Duration(rand.Intn(60)) * time.Second)

	//return security clearance
	c <- crewMapping[name]
}

func Main(){
	rand.Seed(time.Now().UnixNano())

	c1 := make(chan int)
	c2 := make(chan int)

	name := "James"

	go findCrew(name, "Server 1", c1)
	go findCrew(name, "Server 2", c2)

	select {
	case sc := <- c1:
		fmt.Println(name, "has a security clearance of", sc, "found in server1")
	case sc := <- c2:
		fmt.Println(name, "has a security clearance of", sc, "found in server2")
	case <-time.After(25 * time.Second) :
		fmt.Println("Search time out!!")
	}
}