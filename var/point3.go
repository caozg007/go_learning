/*
@Time : 2019-01-16 14:22 
@Author : caozg
@File : point3
*/
package main

import "fmt"

func main() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	/* ptr 不是空指针 */
	if (ip != nil) {
		fmt.Printf("%d 不是空指针\n", ip)
	}
	/* ptr 是空指针 */
	if (ip == nil) {
		fmt.Printf("%d 是空指针\n", ip)
	}

}
