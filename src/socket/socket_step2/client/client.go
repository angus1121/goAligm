package main

import (
	"fmt"
	"net"
	"os"
)

func client(ip string) (err error){

	addr,err:=net.ResolveTCPAddr("tcp4",ip)
	if nil!=err{
		fmt.Println("ResolveTCPAddr err")
		os.Exit(0)
	}

	var conn *net.TCPConn
	conn,err=net.DialTCP("tcp4",nil,addr)
	if nil!=err{
		fmt.Println("Dial failed")
		os.Exit(0)
	}

	defer  conn.Close()

	var n int
	n,err=conn.Write([]byte("this is client"))
	if nil!=err{
		fmt.Println("conn read")
		return nil
	}

	fmt.Println("send mes success",n)
	return nil
}

func main(){
	_=client("127.0.0.1:18013")
}
