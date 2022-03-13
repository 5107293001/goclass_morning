package goroutineexample

import (
	"fmt"
	"time"
)

func initChannelWithSelect() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Millisecond)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(4 * time.Millisecond)
		c2 <- "two"
	}()
	// fmt.Println(<-c1)
	// fmt.Println(<-c2)
	for i := 0; i < 2; i++{
		select {
		case msg1 := <-c1:
		fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case <- time.After(time.Millisecond* 4):
			fmt.Println("timeout")
		// default:
		// 	fmt.Println("default")

		}
	}

}
