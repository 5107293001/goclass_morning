package numberparsing

import (
	"fmt"
	"math/rand"
)

func initRandomNumbers() {

	fmt.Println("======random number=========")
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(1700))
	fmt.Println(rand.Intn(55))
	fmt.Println(rand.Intn(1050))

	fmt.Println(rand.Float64()*10 +5)
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64()+5)
	fmt.Println(rand.Float64())

	

}