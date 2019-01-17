/*
@Time : 2019-01-17 11:09 
@Author : caozg
@File : 多重赋值&变量交换
*/
package main

import "fmt"

func main() {
	a,b,c := 10,20,30
	fmt.Printf("a= %d, b= %d, c=%d \n",a,b,c)

	//交换两个变量的值,传统方法
	x :=33
	y := 44
	var tmp int
	tmp = y
	y = x
	x =tmp
	fmt.Printf("x= %d, y= %d \n",x,y)

	//交换两个变量,go方法
	//i :=10
	//j :=20
	//更简单的声明方式
	i,j :=10,20
	i,j =j,i
	fmt.Printf("i= %d, j= %d \n",i,j)



}
