package main

import (
	"fmt"
	"net/http"
)

func handler(witer http.ResponseWriter,request *http.Request)  {
	fmt.Fprintf(witer,"hello world, %s!", request.URL.Path[1:])
}

func main()  {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}