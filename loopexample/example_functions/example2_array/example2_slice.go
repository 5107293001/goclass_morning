package example2array

import "fmt"

func initSliceExample() {

	s := make([]int, 3)
	fmt.Println(s)
	s [0] = 10
	s [1] = 20
	s [2] = 30
	s = append ( s,40)
	s = append(s, 50)
	fmt.Println("slice s = ", s)
	fmt.Println("lenght of s =", len(s))
	
	c := make([]int, len(s))
	copy (c, s)
	fmt.Println("slice c = ", c)

	slice := s[2:5]
	fmt.Println("slice 1 =", slice)

	slice = s [:2] //0 to 2-1
	fmt.Println("slice 2 =", slice)

	
	slice = s [3:] //3,4

	fmt.Println("slice 2 =", slice)

	
	slice = s [:]

	fmt.Println("slice 2 =", slice)
     ex2 := [] string {"v", "c"}
    fmt.Println(ex2)

//two D with slices (homework)
twoD := make([] [] int,3)

for i:= 0; i < 3; i++ {  
	ilen := i+1
	twoD[i] = make([]int, ilen )
	for j := 0; j < ilen; j++ {
		twoD[i][j] = i+j
	}
}
 fmt.Println("2d slices = ", twoD)

}


 