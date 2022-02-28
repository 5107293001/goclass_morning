package numberparsing

import (
	"fmt"
	"strconv"
)

func InitNumberParsing() {

	testvalue := "100" //base 2 == binary (1,0); decimal == base 10(0-9)
	fmt.Println(testvalue)
	fmt.Printf("type of value = %T\n", testvalue)
	value, err := strconv.ParseInt(testvalue, 2, 64)
	if err != nil{

		fmt.Println(err)
	}
	fmt.Println(value)
	fmt.Printf("type of value = %T\n", value)

	
	testvalue = "1.1"
	fmt.Println(testvalue)
	fmt.Printf("type of value = %T\n", testvalue)
	value1, err := strconv.ParseFloat(testvalue,  64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value1)
	fmt.Printf("type of value = %T\n", value)


    v1, _ :=	strconv.Atoi("11")
	fmt.Println(v1)
	initRandomNumbers()
}
