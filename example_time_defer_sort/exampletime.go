package example_time_defer_sort
import (
	"fmt"
	"time"
)

func InitExampleTime() {

	ticker := time.NewTicker( time.Second)
	done := make (chan bool)
   go func( d chan bool) {
	   for {
		   select{
		   case <- d :
			return
		   case t := <- ticker.C:
			fmt.Println("time = ",t)
		   }
	   }
	   
   }(done)
   time.Sleep(time.Minute)
   ticker.Stop()
   done <- true
   fmt.Println("program stop")

}