package goroutineexample

import (
	"fmt"
	"time"
)

func InitGoroutingExample() {

	routine ("init routine example")

	go routine("goroutine")
  go func () {
   fmt.Println("annomyous function")
	} ()

	go routine("testing")
	time.Sleep(time.Microsecond*10)
	fmt.Println("done")

}

func routine(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, " ", i)
	}
	initChannelExample ()
}