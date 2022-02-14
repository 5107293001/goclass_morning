package examplefunctions

import "fmt"

func InitExampleFunctions() {
  sum ()

}

func sum() {
	var a, b int = 5,10
	sum := a + b
	fmt.Println("sum = ", sum)
}
