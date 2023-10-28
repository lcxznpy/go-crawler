package main

import (
	"fmt"
	rpcdemo "go-crawler/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")

	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result)
	if err != nil {
		fmt.Printf("rpc call error : %v", err)
		return
	}
	fmt.Println(result)
}
