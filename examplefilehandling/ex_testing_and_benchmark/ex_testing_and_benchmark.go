package extestingandbenchmark

import (
	"errors"
	"fmt"
)

func InitTestingAndBenchmark() {
}

func Sum(a, b int) (int, error) {

	if a < 0 || b < 0 {
		return 0, errors.New("a and b must be positive")

	}
	return a + b, nil
}
func Loop() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			fmt.Println("testing")
		}
	}
}
