/*
@Time : 2019-01-30 17:14
@Author : caozg
@File : thread
*/
package main

import (
	"fmt"
	"time"
)

func loop() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("第 %d 个main进程在运行 \n", i)
		time.Sleep(time.Second / 10)
	}
}
func loop2() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("第 %d 个一号go进程在运行 \n", i)
		time.Sleep(time.Second / 10)
	}
}

func loop3() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("第 %d 个二号go进程在运行 \n", i)
		time.Sleep(time.Second / 10)
	}
}

func main() {
	go loop3()
	go loop2()
	loop()
	time.Sleep(time.Second * 50)
}
