package example_time_defer_sort

import (
	"fmt"
	"sort"
)

func InitExampleSort() {
	//sorting

	strs := []string{"d", "a", "c", "b"}

	fmt.Println("string=", strs)

	sort.Strings(strs)

	fmt.Println("sorted string=", strs)
	//panic ("error at mid of the program")
	ints := []int{5, 2, 1, 7, 4}

	fmt.Println("Integers = ", ints)

	s := sort.IntsAreSorted(ints)

	if !s {
		sort.Ints(ints)

		fmt.Println("sorted integers=", ints)

	}

}
