package examplefunctions

import "fmt"

func InitExampleClosures() {
    //aninymous functions
	bar := func() {
		fmt.Println("=================")

	}
	fmt.Println("helllo")
	bar ()
	fmt.Println("world")

	seq := intSeq()
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
     fmt.Println(seq())
    bar()

		seq1 := intSeq()
	fmt.Println(seq1())
	fmt.Println(seq1())

         func () {   //aninymous functions

    fmt.Println("i am inside anonymous functions")
		 } ()

}
   

   func intSeq () func () int {
     i := 0
	 return func () int {
		 i++
		 return i
	 }
   }
