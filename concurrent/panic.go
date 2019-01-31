/*
@Time : 2019-01-30 18:50
@Author : caozg
@File : panic
*/
package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	cha1 := make(chan int)
	go sendData(cha1)
	fmt.Println(<-cha1)
}
