package main

import (
	// "github.com/shres/goclass_morning/examplefilehandling"
	// numberparsing "github.com/shres/goclass_morning/number_parsing"
	// "fmt"
	// "os"

	example_string_functions"github.com/shres/goclass_morning/example_string_functions"
	_ "github.com/shres/goclass_morning/example_time_defer_sort"
	_ "github.com/shres/goclass_morning/number_parsing"

	_ "github.com/shres/goclass_morning/examplefilehandling/goroutineexample"
	_ "github.com/shres/goclass_morning/http_client_server"
	_ "github.com/shres/goclass_morning/number_parsing/example_error_handling"
)

func main() {
	//initials.Init()
	//   loopexample.InitLoopExample()
	//examplefunctions.InitExampleFunctions()
	//examplefunctions.InitSwap()
	// example2array.Init()

	// exampledatastructure.InitDataStructure()
	// examplefunctions.InitVariadicFunction()
	// examplestructs.InitExampleStruct()
	// examplefunctions.InitExampleClosures ()
	// examplestructs.InitExampleStructs2()
	// numberparsing.InitNumberParsing()
	// 	numberparsing.InitRandomNumbers()
	//    examplefilehandling.InitFileHandling()
	//    examplefilehandling.InitFileWrite()
	//    examplefilehandling.InitFileWrite()

	//    extestingandbenchmark.InitTestingAndBenchmark()

	// args := os.Args
	// argsWithoutname := args[1:]
	// fmt.Println(args)
	// fmt.Println(argsWithoutname)
	//   exampleerrorhandling.InitExampleErrorHandling()
	// goroutineexample.InitGoroutingExample()
	// httpclientserver.InitHttpClient()
	// httpclientserver.InitHttpServer()
	// exampletime.InitExampleTime()
	// example_time_defer_sort.InitDeferExample()
	example_string_functions.InitExampleStringFunctions()
}
