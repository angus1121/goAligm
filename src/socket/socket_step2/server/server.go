package main

import (
	"fmt"
	"net"
	"io"
	"os"
)

func listen(conn net.Conn){

	str:=make([]byte,50)

	defer conn.Close()

	for ; ;  {
		n,err:=conn.Read(str)
		if nil != err{
			if err == io.EOF {
				fmt.Println("remote sockt closed")
				return
			}
			fmt.Println("receive err")
			continue
		}
	fmt.Println("receive mes:",string(str),"length",n)
	}
	return
}

func server(ip string,port int) (err error) {
	add :=fmt.Sprintf("%s:%d",ip,port)
	addr,err:= net.ResolveTCPAddr("tcp",add)
	fmt.Println(addr.String())

	listener,err:=net.Listen("tcp",addr.String())
	defer listener.Close()
	if nil != err{
		fmt.Println("DialTCP failed")
		os.Exit(0)
		return nil
	}

	for ; ;  {
		conn,err:=listener.Accept()
		if nil != err{
			fmt.Println("Accept failed")
			continue
		}

		go listen(conn)
	}

	return nil
}

func main(){
	_=server("127.0.0.1",18013)
}
