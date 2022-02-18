package exampledatastructure

import "fmt"

func initExampleMap() {

	m := make(map[string]int) //map [key-data-type]value-data-type

	// "a" => 97
	// "b" => 98
	// "c" => 99
	m["a"] = 97
	m["b"] = 98
	m["c"] = 99
	fmt.Println("map=",m)

	v1, ok := m["a"]
	fmt.Println(ok)
	if ok {
		fmt.Println("value of a =",v1)
	}

	m ["d"] = 100
	fmt.Println("map=",m)
fmt.Println("length of map=",len(m))
delete(m,"c")
fmt.Println("map=",m)



//exampleRange()

}
