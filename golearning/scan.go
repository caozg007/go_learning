/*
@Time : 2019-01-28 14:02
@Author : caozg
@File : scan
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var quote string
	var name string

	fmt.Print("What is the quote? ")
	fmt.Scanf("%s", &quote)

	fmt.Print("Who said it? ")
	fmt.Scanf("%s", &name)

	fmt.Printf("%s says, \"%s\"", strings.Title(name), quote)
}
