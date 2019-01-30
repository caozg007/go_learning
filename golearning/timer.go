/*
@Time : 2019-01-28 10:54
@Author : caozg
@File : function
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	//启动第一个线程，id为1
	go task(1)
	time.Sleep(time.Second * 2)
	//启动第二个线程，id为2
	go task(2)
	//time.Sleep(time.Second * 10)
	var quit string
	fmt.Scanf("%s", &quit)
}

func task(id int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("第%d个线程运行,打印i的值为：%d \n", id, i)
		time.Sleep(time.Second)
	}
}
