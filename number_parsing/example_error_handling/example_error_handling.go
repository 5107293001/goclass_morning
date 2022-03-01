package exampleerrorhandling

import (
	"errors"
	"fmt"
)

func InitExampleErrorHandling() {

   n, err :=	fmt.Println("error example")

   Panicerrorhandle(err)

   fmt.Println("value of n =",  n)
   pNo, err := positiveno(10)
  
   Panicerrorhandle(err)
   fmt.Println("positiveno = ", pNo)

   pNo, err = positiveno(-1)
  
   Panicerrorhandle(err)
   fmt.Println("positiveno = ", pNo)
}
func Panicerrorhandle(err error) {
	if err != nil {
		panic(err)
	}
}
func positiveno(arg int) (int, error) {
	if arg < 0{
		return 0, errors.New("args can not be negative")
	}
  return arg, nil

}