package reflection

import (
	"reflect"
	"fmt"
)

func ReflectionOnStruct() {

	type myStruct struct {
		Field1 int `alias:"f1" desc:"field number 1"`
		Field2 string `alias:"f2" desc:"field number 2"`
		Field3 float64 `alias:"f3" desc:"field number 3"`
	}

	mys := myStruct{2, "Hello", 2.4}

	inspectStructType(&mys)

}

func inspectStructType(i interface{}) {

	mysRValue := reflect.ValueOf(i)
	if mysRValue.Kind() != reflect.Ptr {
		return
	}
	mysRValue = mysRValue.Elem()
	if mysRValue.Kind() != reflect.Struct {
		return
	}

	mysRValue.Field(0).SetInt(15)

	mysRType := mysRValue.Type()

	for i:= 0; i < mysRType.NumField(); i++ {

		fieldRValue := mysRValue.Field(i)
		fieldRType := mysRType.Field(i) // data type: StructField

		// to use fieldRValue.Interface() methods should be exportable

		fmt.Printf("Field Name: '%s', field type: '%s', field value: '%v' \n", fieldRType.Name, fieldRType.Type, fieldRValue.Interface())
		fmt.Println("Struct tags, alias:", fieldRType.Tag.Get("alias"), "description: ", fieldRType.Tag.Get("desc"))

	}
}