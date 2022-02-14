package loopexample

import "fmt"

func InitLoopExample() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		
	}
	 // i = 0; i < 10; 
    // i = 2
    //..
   // i = 10
	fmt.Println("sum=", sum)

	sum = 1
	for sum < 45 {
		sum += sum
		fmt.Println("sum In loop=", sum)
		fmt.Println(sum < 45)
	}
		fmt.Println("sum2 =", sum)
		a := 0
		for a < 10 {
			if a == 5 {
				/* skip the iteration */ a = a + 1
				continue

			} 
			fmt.Printf("value of a: %d\n",a)
			a++
		}
		fmt.Println("============")
		fmt.Println("Using break")
      a = 0
	for a  < 10 {
		fmt.Printf("value of a: %d\n", a)
		a++
		if a > 5 {
			break
		}
	}
}