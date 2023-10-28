package main

import (
	"fmt"
	rpcdemo "go-crawler/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//注册服务
	rpc.Register(rpcdemo.DemoService{})
	//监听端口
	l, err := net.Listen("tcp", ":1234")
	fmt.Println("wait for accept")
	if err != nil {
		panic(err)
	}
	for {
		//等待连接
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accecpt error : %v", err)
			return
		}
		fmt.Println("conn success create")
		//使用连接
		go jsonrpc.ServeConn(conn)
	}
}
