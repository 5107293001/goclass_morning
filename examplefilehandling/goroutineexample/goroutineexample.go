package goroutineexample

import "fmt"

func InitGoroutingExample() {

}

func routine(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, " ", i)
	}
}