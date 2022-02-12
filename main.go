package main

import "fmt"
const PI = 3.14 //constant global variable
var i,j,k  = 1, 2.2 ,"abc" //global variable

func main() {

fmt.Println("Hello go")
	var age uint8 = 100 //memory 8 bit= 1 byte
	var B float32 = 1.1// local variacle
	fmt.Println("value of a=", age)
	fmt.Println(B)
	i =2
	fmt.Println(i,j,k)
	y := 10.10
	fmt.Printf("Type of y = %T", y)

	example()
}
func example() {

	var age uint8 = 80
	fmt.Println("In example =",age,PI)
}