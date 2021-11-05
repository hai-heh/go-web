package main

import (
	"bufio"
	"fmt"
	"net"
)

func main()  {

	listen,err:=net.Listen("tcp","localhost:20000")
	if(err!=nil){
		fmt.Println("err:",err)
		return
	}
	for{
		coon,err:=listen.Accept()
		if(err!=nil){
			fmt.Println("err:",err)
			return
		}
		go slove(coon)
	}
}

func slove(conn net.Conn) {

	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if (err != nil) {
			fmt.Println("err:", err)
			return
		}
		s := string(buf[:n])
		fmt.Println(s)
		conn.Write([]byte(s))
	}

}
