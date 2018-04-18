package main

import (
	"fmt"
	"net"
	"os"
)

func client() {
	//fmt.Println("client ip %s port %d",addr,port)
	conn, err :=net.Dial("tcp","127.0.0.1:18013")
	if err != nil{
		fmt.Println("connect failed")
		os.Exit(0)
	}
	defer conn.Close()

	conn.Write([]byte("hello server"))

	println("send msg")
}

func main(){
	client()
}