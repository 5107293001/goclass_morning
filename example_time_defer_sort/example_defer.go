package example_time_defer_sort

import "fmt"

func InitDeferExample() {
	defer fmt.Println("stopping the program 1")
	defer fmt.Println("stopping the program 2")
	defer fmt.Println("stopping the program 3")
	a := 10
	b := 13
	fmt.Println(a, b)
	testDefer ()
	fmt.Println("Hello Golang")
}

func testDefer () {
	defer fmt.Println("stopping function test defer 1")
	defer fmt.Println("stopping function test defer 2")
	defer fmt.Println("stopping function test defer 3")
	fmt.Println("running test defer function")
	fmt.Println("3 * 2 = ", 3*2)
}
