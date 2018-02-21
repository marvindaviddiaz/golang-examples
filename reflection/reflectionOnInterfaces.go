package reflection

import (
	"reflect"
	"fmt"
)

type Printer interface {
	Print(s string)
}

type printerStruct struct {
	s string
}
func (p *printerStruct) Print(s string) {
	p.s = s
	fmt.Println(s)
}

func ReflectionOnInterfaces() {
	p := new(printerStruct)
	inspectType(p)
}

func inspectType(obj interface{}) {
	v := reflect.ValueOf(obj)
	t := v.Type()

	myInterface := reflect.TypeOf((*Printer)(nil)).Elem() // trick to get the Type of an interface (empty pointer)

	fmt.Println("obj implements Printer?",  t.Implements(myInterface))

	if t.Implements(myInterface) {
		printFunc := v.MethodByName("Print")
		args := [] reflect.Value { reflect.ValueOf("Printing Hello from a reflection object!!!") }
		printFunc.Call(args)
	}


}