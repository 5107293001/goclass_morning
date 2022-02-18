package examplefunctions

import "fmt"
func InitSwap() {
    var x, y int = 10, 20

    fmt.Println("=====Swap By Value=====")
    fmt.Printf("IN Main Before Swap::Value of x = %d, Value of y = %d\n", x, y)
    swapbyvalue(x, y)
    fmt.Printf("IN Main After Swap::Value of x = %d, Value of y = %d\n", x, y)

    fmt.Println("=====Swap By Reference=====")

    fmt.Printf("IN Main Before Swap::Value of x = %d, Value of y = %d\n", x, y)
    swapbyreference(&x, &y)
    fmt.Printf("IN Main After Swap::Value of x = %d, Value of y = %d\n", x, y)
}

func swapbyvalue(a, b int) int {
    fmt.Printf("IN SWAP Before Swap::Value of x = %d, Value of y = %d\n", a, b)
    temp := a
    a = b
    b = temp
    fmt.Printf("IN SWAP After Swap::Value of x = %d, Value of y = %d\n", a, b)
    return temp
}

func swapbyreference(a, b *int) int {
    fmt.Printf("IN SWAP Before Swap::Value of x = %d, Value of y = %d\n", *a, *b)
    var temp int
    temp = *a
    *a = *b
    *b = temp
    fmt.Printf("IN SWAP After Swap::Value of x = %d, Value of y = %d\n", *a, *b)
    return temp
}

