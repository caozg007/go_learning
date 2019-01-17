/*
@Time : 2019-01-16 14:12 
@Author : caozg
@File : swap
*/
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Mahesh", "Kumar")
	fmt.Println(b, a)
	fmt.Println(swap("Mahesh", "Kumar"))

}