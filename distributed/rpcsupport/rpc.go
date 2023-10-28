package rpcsupport

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 创建rpc服务端，设置监听端口和要注册的服务
func ServerRpc(host string, service interface{}) error {
	//注册服务
	rpc.Register(service)
	//监听端口
	l, err := net.Listen("tcp", host)
	fmt.Println("wait for accept")
	if err != nil {
		return err
	}
	for {
		//等待连接
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accecpt error : %v", err)
			continue
		}
		fmt.Println("conn success create")
		//使用连接
		go jsonrpc.ServeConn(conn)
	}
}

// 根据host建立连接
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)

	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
