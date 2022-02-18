package example2array

import "fmt"
func Init () {

	//initArrayExample()
	initSliceExample()
}

func initArrayExample() {

	var arrayExample [3]int
	arrayExample[0] = 10
	arrayExample[1] = 11
	arrayExample[2] = 12
	fmt.Println(arrayExample)

	arrayExample2 := [5] int {5,10,15,20,25}
	fmt.Println(arrayExample2)

	for i := 0; i < len(arrayExample2); i ++ {
        fmt.Printf("index = %d and value = %d\n ", i, arrayExample2)

	}

	arrayExample3 := [2] string {"rajan ", "rupa"} 
	fmt.Println(arrayExample3)

	var arrayExample4 [2] [3]int 
	for i := 0; i < 2; i ++ {
		for j := 0; j < 3; j++ {
			arrayExample4[i][j] = i + j
			fmt.Println(arrayExample4[i][j])
		}
	}
	fmt.Println("Two D array =", arrayExample4)
}