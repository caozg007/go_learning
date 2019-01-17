/*
@Time : 2019-01-15 18:24 
@Author : caozg
@File : mutilvar
*/
package main

var x, y int

var ( // 这种因式分解关键字的写法一般用于声明全局变量
	A int = 2
	B bool = true
)

var C, D int = 1, 2
var E, F = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//G, H := 123, "hello"

func main() {
	G, H := 123, "hello"
	println(x, y, A, B, C, D, E, F, G, H)
	println(&x, &y, &A, &B, &C, &D, &E, &F, &G, &H)
}
// 0x10d4b30 0x10d4b38 0x10b6140 0x10b6120 0x10b6148 0x10b6150 0x10b6158 0x10b6da0 0xc000042770 0xc000042778