package errors

import (
	"time"
	"math/rand"
	"fmt"
)

type NotFoundException struct {
	Name, Server, Msg string
}
func (e NotFoundException) Error() string {
	return e.Msg
}

var crewMapping = map[string]int {
	"James": 5,
	"Kevin": 10,
	"Rahul": 9,
}

func findCrew(name, server string) (int, error) {
	if v, ok :=  crewMapping[name]; !ok {
		//return -1, errors.New("Crew member not found")
		//return -1, fmt.Errorf("Crew member '%s' could be not found on server '%s' ", name, server)
		return -1, NotFoundException{name, server, "Crew member not found"}
	} else {
		return v, nil
	}
}

func Main() {
	rand.Seed(time.Now().UnixNano())
	if clearance, err := findCrew("Ruko", "Server 1"); err != nil {

		if v, ok := err.(NotFoundException); ok { // instance of NotFoundException
			fmt.Printf("Crew member '%s' could be not found on server '%s' ", v.Name, v.Server)
		} else {
			fmt.Println("Error ocurred:", err)
		}

	} else {
		fmt.Println("Clearance level found:", clearance)
	}
}