/*
@Time : 2019-01-30 18:17
@Author : caozg
@File : channel
*/
package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T \n", a)
		fmt.Println(a)
	}
}
