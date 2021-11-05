package main

import "net/rpc"

type register struct {

}

func (this *register) registerHelloService(serverName string,Interface HelloServerInterface)  {
	rpc.RegisterName(serverName,Interface)
}