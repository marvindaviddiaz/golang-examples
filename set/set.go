package set

import "fmt"

type Set map[string]struct{} // declaring a type

func Run() {
	s := make(Set) // map initialization
	s["item1"] = struct{}{}
	s["item2"] = struct{}{}

	fmt.Println(getSetValues(s))
}

func getSetValues(s Set) []string {
	var retVal []string
	for k, _ := range s {
		retVal = append(retVal, k)
	}
	return retVal
}
