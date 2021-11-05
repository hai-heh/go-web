package main

import (
	"bufio"
	"fmt"
	"net"
)

func main()  {

	conn, err := net.Dial("tcp", "localhost:20000")
	if(err!=nil){
		fmt.Println(err)
		return
	}
	for{
		var s string
		fmt.Scanln(&s)
		_, err := conn.Write([]byte(s))
		if(err!=nil){
			return
		}
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if(err!=nil){
			return
		}
		fmt.Println(string(buf[:n]))
	}

}