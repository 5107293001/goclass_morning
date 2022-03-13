package httpclientserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func InitHttpServer() {

	fmt.Println("String Web Server")
   http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request){
	   rw.Write([]byte("Hello Rajan Shrestha"))
   })

   http.HandleFunc("/hello",Hello)
http.ListenAndServe(":8090",nil)
}
func Hello(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		fmt.Println(r.URL.Query()["id"])
		rw.Write([] byte(`{"id" :1, "name" : "Rupa"}`))
		
	} else if r.Method == "POST" {

		fmt.Println(r.Header)
		body , err := ioutil.ReadAll(r.Body)
		if err != nil{
			panic((err))
		}
		fmt.Println(string(body))
		rw.Write([]byte(`{"message: "succussfully created data"}`))

	
	}else if r.Method == "DELETE" {
		rw.Write([]byte(`{"message: "succussfully deleted data"}`))
		

	}else if r.Method == "PUT" {
		rw.Write([]byte(`{"message: "succussfully updated data"}`))
	}
	
}