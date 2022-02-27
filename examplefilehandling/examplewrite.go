package examplefilehandling

import (
	"fmt"
	"os"
)

func InitFileWrite() {

	data := []byte("golang!\n Hello there\n")
	err := os.WriteFile("example.txt",data, 0644)
	p(err)
	defer file.close()
	n1, err := file.write(data)
	p(err)
	fmt.Printf("data written =%d\n",n1)
	
}