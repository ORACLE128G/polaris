package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"polaris/rpc/server"
)

func main() {
	// Call json rpc method.
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div",
		rpcserver.Args{A: 10, B: 5},
		&result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("division result: %v", result)

}
