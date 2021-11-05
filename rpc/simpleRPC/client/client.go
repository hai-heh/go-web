package main

import (
	"fmt"
	"net/rpc"
)

const serverName string  = "helloService"

type helloServiceClient struct {
	client *rpc.Client
}

func (this *helloServiceClient) getServer(address string)  {
	client,e:= rpc.Dial("tcp", address)
	if(e!=nil){
		fmt.Println(e)
		return
	}
	this.client=client
}

func (this *helloServiceClient) Hello(methodName string,request string,reply *string) error {
	this.client.Call(serverName+"."+methodName,request,reply)
	return nil
}



func main()  {
	var client helloServiceClient
	client.getServer("localhost:1234")
	var reply string
	client.Hello("Hello","hello",&reply)
	fmt.Println(reply)
}