/*
@Time : 2019-01-30 18:29
@Author : caozg
@File : channel3.go
*/
package main

import (
	"fmt"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
	fmt.Println("calcSquares")
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		fmt.Printf("%d \n", digit)
		sum += digit * digit * digit
		fmt.Printf("%d \n", sum)
		number /= 10
	}
	cubeop <- sum
	fmt.Println("calcCubes")
}

func main() {
	number := 125
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
	//squares := <-sqrch
	//fmt.Println("Final output", squares)
}
