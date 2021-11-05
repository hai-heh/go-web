package myProto

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "hello"})
	fmt.Println(reply.GetValue())
	stream, err := client.Server(context.Background(), &String{Value: "hello"})
	for{
		recv, err := stream.Recv()
		if(err==io.EOF){
			break;
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(recv)
	}
}


// 生成grpc代码
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=../proto/ hello.proto hello.message.proto
