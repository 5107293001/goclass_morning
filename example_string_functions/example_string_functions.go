package example_string_functions

import (
	"fmt"
	"strings"
)

var print = fmt.Println

func InitExampleStringFunctions() {

	print("contains :", strings.Contains("golang", "lang"))
	print("Hasprefix :", strings.HasPrefix("golang", "gol"))
	print("count :", strings.Count("golnag", "g"))
	print("HasSuffix :", strings.HasSuffix("golang", "ng"))
	print("Index :", strings.Index("golang", "la"))
	print("repeat :", strings.Repeat("go ",  10))
	print("replace: ", strings.Replace("golang!", "g","$",-1))
	print("split:", strings.Split("g-o-l-a-n-g", "-"))
	print("split:", strings.Split("go lang", " "))
	print("join:", strings.Join([]string{"g","o","l","a","n","g"},"-"))
	print("join:", strings.Join([]string{"go", "lang"},""))
	print("toupper:", strings.ToUpper("GoLang"))
	print("tolower:", strings.ToLower("GoLang"))

	fmt.Printf("boolean\t\tvalue = %t\ninterger\t\t value =%d\n", true,1)
  //formatted printf
  fmt.Printf("int =%d\n", 2)
	fmt.Printf("binary = %b\n", 4)
	fmt.Printf("char = %c\n", 97)
	fmt.Printf("float value = %f\n", 1.1)
	fmt.Printf("string value = %s\n", "golang")
 l:= location{
	 lat : 21.05415,
	 lng : 22.55100,
 }
 fmt.Printf("struct 1 = %v\n", l)
 fmt.Printf("struct 2 = %+v\n", l)
 fmt.Printf("struct 3 = %#v\n", l)
 fmt.Printf("|%10s|%10s|\n", "min price", "max price")
fmt.Printf("|%10d |%10d|\n ", 10,12)
fmt.Printf("|%10d| %10d|\n ", 10000,12000)
fmt.Printf("|%10d| %10d|\n ", 1,15000)

fmt.Printf("|%-10d |%-10d|\n ", 10,12)
fmt.Printf("|%-10d| %-10d|\n ", 10000,12000)
fmt.Printf("|%-10d| %-10d|\n ", 1,15000)
s := fmt.Sprintf("value of a = %d\n", 15 )
fmt.Println(s)
}
type location struct {
	lat , lng float64
}
