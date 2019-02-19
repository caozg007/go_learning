/*
@Time : 2019-02-19 15:56
@Author : caozg
@File : reflecttest
*/
package main

import (
	"fmt"
	"reflect"
)

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "caozg",
		Ice:  true,
	}
	b := Drink{
		Name: "caozg",
		Ice:  true,
	}
	fmt.Printf("%s\n", reflect.TypeOf(a))
	fmt.Printf("%s", reflect.TypeOf(b))
}
