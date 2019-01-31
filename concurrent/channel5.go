/*
@Time : 2019-01-30 18:59
@Author : caozg
@File : channel5
*/
package main

import (
	"fmt"
)

func producer2(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
func main() {
	ch := make(chan int)
	go producer2(ch)
	for v := range ch {
		fmt.Println("Received ", v)
	}
}
