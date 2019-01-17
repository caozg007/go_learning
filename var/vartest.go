/*
@Time : 2019-01-15 18:22 
@Author : caozg
@File : vartest
*/
package main

var a = "菜鸟教程"
var b string = "runoob.com"
// 布尔类型bool
var c bool

func main() {
	println(a, b, c)

	//自动推导类型，只能使用一次，用于变量初始化，相当用户var int a ; a =30
	//s := 30
	var  s int
	s = 30
	println(s)
	//已声明变量s,不能再次声明
	//s :=40
	//println(s)


}
