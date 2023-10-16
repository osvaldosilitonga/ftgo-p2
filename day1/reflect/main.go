package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number float64 = 23.42
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variable :", reflectValue)

	if reflectValue.Kind() == reflect.Float64 {
		fmt.Println("nilai variable :", reflectValue.Float())
	}

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variable :", reflectValue.Int())
	}
}
