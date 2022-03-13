package goroutineexample



func initChannelExample2() {

// 	done := make(chan bool, 2)
// 	fmt.Println("starting main thread ......")
// 	go worker(done)
// 	<- done
	
// }

// func worker(done chan bool) {
// 	fmt.Println("working on thread ......")
// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j <10; j++{
// 			for k := 0; k < 1000; k++ {
// 				sum = sum +k
// 			}
			
// 		}
// 		fmt.Println("value of i =", i ) 
			
// 		}
// 		done <- true
		initChannelWithDirection()
		initChannelWithSelect()
	}


