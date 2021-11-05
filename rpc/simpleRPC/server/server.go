package main

import (
	"log"
	"net"
	"net/rpc"
)

const serverName string  = "helloService"


type HelloServerInterface interface {
	Hello(request string,reply *string) error
}

type HelloServer struct {

}


func (p *HelloServer) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	r:=register{}
	r.registerHelloService(serverName,new(HelloServer))
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}