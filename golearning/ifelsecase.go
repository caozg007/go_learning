/*
@Time : 2019-01-28 10:54
@Author : caozg
@File : function
*/

package main

import "fmt"

func main() {
	x := 180

	if x > 200 {
		println("大于200")
	} else if x <= 100 {
		fmt.Println("小于等于100")
	} else {
		fmt.Println("100和200之间")
	}

	switch {
	case x > 200:
		fmt.Println("CASE:大于200")
	case x < 100:
		fmt.Println("CASE:小于等于100")
	default:
		fmt.Println("CASE:100和200之间")
	}

	for y := 0; y < 10; y++ {
		fmt.Printf("%d \t ", y)
	}

	for z := 5; z >= 0; z-- {
		fmt.Printf("\n %d  ", z)
	}
	fmt.Println("\n")
	s := []int{100, 200, 300}
	for m, n := range s {
		fmt.Println(m, ":", n)
	}
}
