package appliances

import (
	"errors"
	"fmt"
)

type Appliance interface {
	Start()
	GetPurpose() string
}

const (
	STOVE = iota // iota increments value to all enum values starting by 0
	FRIDGE
	MICROWAVE
)

func CreateAppliance(mtype int) (Appliance, error) {
	switch mtype {
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
	case MICROWAVE:
		return new(Microwave), nil
	default:
		return nil, errors.New("Invalid Appliance Type")
	}
}


func Main() {
	fmt.Println("Enter preferred appliance type")
	fmt.Println("0: Stove")
	fmt.Println("1: Fridge")
	fmt.Println("2: Microwave")

	var myType int
	fmt.Scan(&myType)

	myAppliance, err := CreateAppliance(myType)

	if err == nil {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		fmt.Println(err)
	}
}