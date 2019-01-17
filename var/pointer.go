/*
@Time : 2019-01-16 09:24 
@Author : caozg
@File : pointer
*
*/
package main
import "fmt"
func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	//var intP *int
	//intP = &i1
	var intP = &i1
	// *intP , intP为5的存储地址：0xc00007a008 ， *内存地址，代表：0xc00007a008作为地址的那个空间的取值
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	//打印intP的数据类型
	fmt.Printf("%T\n",intP)

	var p = new(int)
	*p = 666
	fmt.Println("*p = ", *p)
	fmt.Println("p = ", p)
	//打印P的数据类型 *int
	fmt.Printf("%T\n",p)
	var s int =10
	//打印s的数据类型 int
	fmt.Printf("%T\n",s)

}