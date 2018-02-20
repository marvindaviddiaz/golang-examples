package embedding

import (
	"math/rand"
	"time"
	"fmt"
)

type customRand struct {
	*rand.Rand
	count int
}

func NewCustomRand(i int64) *customRand {
	return &customRand {
		Rand: rand.New(rand.NewSource(i)),
		count: 0,
	}
}

func (cr *customRand) RandRange(min, max int) int {
	cr.count++
	return cr.Rand.Intn(max - min) + min
}

func (cr *customRand) GetCount() int {
	return cr.count
}

func Main() {
	cr := NewCustomRand(time.Now().UnixNano())
	fmt.Println("Custom Rand:", cr.RandRange(5, 30))
	fmt.Println("Inherited Rand:", cr.Intn(10)) // The same as cr.Rand.Intn(10)
	fmt.Println("Counter:", cr.GetCount())
}