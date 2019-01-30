/*
@Time : 2019-01-17 15:32
@Author : caozg
@File : iota
*/
package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a=%d,b=%d,c=%d", a, b, c)
}
