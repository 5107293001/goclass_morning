package examplefilehandling

import (
	"fmt"
	"os"
)

func InitFileHandling() {

	output, err := os.ReadFile("test.txt")
	p(err)
	fmt.Println(string(output))
	file,err := os.Open("test.txt")
	p(err)
	b1 := make([]byte,3)
	n1, err := file.Read(b1)
	p(err)
	fmt.Printf("output =%s and n = %d\n",string (b1,n1)
}