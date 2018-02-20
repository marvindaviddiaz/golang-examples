package interfaces

import "fmt"

type Node interface {
	GetValue () int
}

// double
type DoubleNode struct {
	value int
}
func (node *DoubleNode) GetValue() int {
	return node.value * 2
}
// triple
type TripleNode struct {
	value int
}
func (node *TripleNode) GetValue() int {
	return node.value * 3
}

//func getOriginalValue(node interface{}) int { // THIS LINE IS LIKE THE MOST GENERIC WAY
func getOriginalValue(node Node) int {
	// switch with instance of
	switch concrete := node.(type) {
	case *DoubleNode:
		return concrete.GetValue() / 2
	case *TripleNode:
		return concrete.GetValue() / 3
	default:
		return -1
	}
}

func Main() {
	nodeDouble := DoubleNode{12}
	nodeTriple := TripleNode{33}

	fmt.Println("Double Value:",nodeDouble.GetValue())
	fmt.Println("Triple Value:", nodeTriple.GetValue())

	fmt.Println("Double Original Value:", getOriginalValue(&nodeDouble))
	fmt.Println("Triple Original Value:", getOriginalValue(&nodeTriple))


}