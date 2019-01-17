/*
@Time : 2019-01-17 11:35 
@Author : caozg
@File : 匿名变量
*/
package main

import "fmt"

func test() (x,y,z int){
	return 1,2,3
}

func main() {
     x,y,z :=0,0,0
	//var x,y,z=test()
	_,y,_=test()
	fmt.Printf("x= %d,y=%d,z=%d \n", x,y,z)
}
