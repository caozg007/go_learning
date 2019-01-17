/*
@Time : 2019-01-15 18:54 
@Author : caozg
@File : point
*/
package main

import "fmt"

func main() {
	//定义一个变量i等于42
	i := 42
	//方式一
	//&i指针，生成一个指针，指向作用对象i
	p := &i
	//*T符号表示指针指向的底层的值
	fmt.Println(*p)
	//通过这个指针去操作i，去改变i的值
	*p = 21
	fmt.Println(i)
	//定义一个变量j,值为2701
	j := 2701
	//方式二
	//&i指针，生成一个指针，指向作用对象i
	//*p(除后的值) = *p（2701）/37
	p = &j
	*p = *p / 37
	fmt.Println(j)
}
