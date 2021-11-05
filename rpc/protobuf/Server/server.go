package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
    "web/rpc/protobuf/Client"
)



type HelloServerServiceImpl struct {
	myProto.UnimplementedHelloServiceServer
}


func (p *HelloServerServiceImpl) Hello(
	ctx context.Context, args *myProto.String,
) (*myProto.String, error) {
	reply := &myProto.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (p *HelloServerServiceImpl) Channel(stream myProto.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		reply := &myProto.String{Value: "hello:" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func (p *HelloServerServiceImpl) Server(s *myProto.String, stream myProto.HelloService_ServerServer) error {
	for i:=0;i<6;i++{
		stream.Send(s)
	}
	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	myProto.RegisterHelloServiceServer(grpcServer, new(HelloServerServiceImpl))
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}