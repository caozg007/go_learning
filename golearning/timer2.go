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
	checktime()
}

func checktime() {
	dt := time.Date(2020, 1, 10, 0, 0, 1, 100, time.Local)
	fmt.Println(time.Now().After(dt))  // true
	fmt.Println(time.Now().Before(dt)) // false

	dt2 := time.Now()
	fmt.Println(dt2)

	// 是否相等 判断两个时间点是否相等时推荐使用 Equal 函数
	fmt.Println(dt2.Equal(time.Now()))
	fmt.Println(time.Now())
}
