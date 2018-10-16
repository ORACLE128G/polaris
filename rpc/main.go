package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"polaris/rpc/server"
)

func main() {
	rpc.Register(rpcserver.DemoService{})
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
