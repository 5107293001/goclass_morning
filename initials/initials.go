package initials

import (
	"fmt"
	"strings"
)
const PI = 3.14 //constant global variable
var i,j,k  = 1, 2 ,"abc" //global variable

func Init() {

fmt.Println("Hello go")
	var age uint8 = 100 //memory 8 bit= 1 byte
	var B float32 = 1.1// local variacle
	fmt.Println("value of a=", age)
	fmt.Println(B)
	i =2
	i += 2 // i= i+2
	fmt.Println(i,j,k)
	y := 10.10
	fmt.Printf("Type of y = %T", y)

	example()
      example2()
ifElseExample()


}
func ifElseExample() {
	age := 9
	if age <=1 {
		fmt.Println("infant")

	} else if age > 1 && age <= 25 {
		fmt.Println("teen")

	} else if age > 25 && age <= 64 {
		fmt.Println("youth")

	} else {
		fmt.Println("old")
	}

	i := "go" 
	fmt.Println(i) 
	switch  1 {
	case 1:
		fmt.Println("sum = 10")
	case 2:
		fmt.Println("sum = 8")
	case 3:
		fmt.Println("sum = 5")
	default:
		fmt.Println("default")
	}
}
func example() {

	var age uint8 = 80
	fmt.Println("In example =",age,PI)
}
func example2() {
	/*
	000 = 0
	001 = 1
	010 = 2
	011 = 3
	100 = 4
	101 = 5 letf shift 5 << 2 == 01010 = 10
	110 = 6
	111 = 7
	1000 = 8 
	*/

	var a = 2
	a = a << 2

	fmt.Println(a)
	var b, c = 2,3
	d := b & c
	fmt.Println(d)
	var greeting1 = " hello"
	var greeting2 = "world"
	var check bool = true
	 greeting := greeting1 + greeting2
	 greetingUpper := strings.ToUpper(greeting)
	 fmt.Println(greetingUpper)
	 fmt.Println(greeting,check)

	fmt.Printf(greeting)

	var p *int
	i := 42
	p = &i
	fmt.Println("pointer value =", p)
}