/*
@Time : 2019-01-15 17:28 
@Author : caozg
@File : testgo
*/
package main

import "fmt"

func sayHi() (x, y string) {
	x = "hello"
	y = "world"
	return
}
func main() {
	fmt.Println(sayHi())
}
