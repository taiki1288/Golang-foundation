package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	b := 1
	num_bool := a < b

	fmt.Println(num_bool)
	fmt.Println(reflect.TypeOf(num_bool))
}