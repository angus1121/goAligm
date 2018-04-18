package main

import (
	"fmt"
	"net"
	"os"
	"io"
)

func server(conn net.Conn){
	defer conn.Close()
	msg := make([]byte,50)
	for ; ;  {
		n,err := conn.Read(msg)
		if err != nil&& err!=io.EOF{
			fmt.Println("read err ")
			os.Exit(0)
		}
		if err == io.EOF{
			fmt.Println(conn.RemoteAddr().String(),"socket close ")
			os.Exit(0)
		}

		if n > 0 {
			fmt.Println("receive success ",n)
			fmt.Println(string(msg[0:n]))
		}
	}
}
func main() {
	conn,err:=net.Listen("tcp","127.0.0.1:18013")

	if(err != nil){
		fmt.Println("listen err")
	}
	defer  conn.Close()

	for ; ;  {
		new_conn,err:=conn.Accept()
		if err != nil{
			fmt.Println("accept err")
			os.Exit(0)
		}
		go server(new_conn)
	}

}

