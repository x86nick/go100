package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "civic" // civic string literal will be assigned to name
	fmt.Println(reflect.TypeOf(name))
}
