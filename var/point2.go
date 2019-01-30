/*
@Time : 2019-01-15 19:06
@Author : caozg
@File : point2
*/
package main

import (
	"fmt"
)

func main() {
	//定义一个int类型的变量x
	var x int
	//定义一个？？？？
	var x_ptr *int
	//int类型变量x，x的值为10
	x = 10
	// 将变量x的内存地址，如：0xc00000c028 赋值给x_ptr， 即x_ptr=0xc00000c028
	x_ptr = &x
	// 输出内存地址为ox1001
	fmt.Println(&x_ptr)

	var t = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", t, &t)
	var a *int
	a = &t
}
