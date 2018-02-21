package reflection

import (
	"reflect"
	"fmt"
)

type myFloat float64

func Main() {
	var x1 float32 = 5.7

	inspectIfTypeFloat(x1)

	// Custom type

	var x3 myFloat = 5.7
	v := reflect.ValueOf(x3)
	fmt.Println("\nType:", v.Type())
	fmt.Println("Is type float64?", v.Kind() == reflect.Float64)


	// ## THIRD LAW

	var x2 float32 = 5.8
	v2 := reflect.ValueOf(&x2) //using a pointer
	fmt.Println("\nv2 Settable:", v2.CanSet())
	v2pElem := v2.Elem() // x1
	fmt.Println("v2pElem Settable:", v2pElem.CanSet())
	v2pElem.SetFloat(2.2)

	fmt.Println("Setted Value by Reflection:", x2)
}

func inspectIfTypeFloat(i interface{}) {

	// ## FIRST LAW

	v := reflect.ValueOf(i)

	fmt.Println("Type:", v.Type()) // same as reflect.TypeOf(x1)

	fmt.Println("Is type float64?", v.Kind() == reflect.Float64)

	fmt.Println("Float Value:", v.Float()) // this methods return the max capacity for the types (float64)

	// ## SECOND LAW

	interfaceValue := v.Interface() // inverse of reflect.ValueOf(i)
	switch t:= interfaceValue.(type) {
		case float32:
			fmt.Println("Original float32 value", t)
	}
}