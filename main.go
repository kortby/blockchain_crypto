package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/kortby/blockchaincrypto/node"
	"github.com/kortby/blockchaincrypto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	node := node.NewNode()
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	proto.RegisterNodeServer(grpcServer, node)
	fmt.Println("blockchain running on port: ", ":3000")
	go func () {
		for {
			time.Sleep(time.Second * 3)
			makeTransactions() 
		}
	}()
	grpcServer.Serve(ln)

}

func makeTransactions() {
	client, err := grpc.Dial(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	c := proto.NewNodeClient(client)

	version := &proto.Version{
		Version: 32,
		Height: 100,
	}
	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatal(err)
	}
}