package rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type RpcServer struct {
	Status bool
	Addr string
	handle interface{}

}
func NewRpcServer(addr string,handle interface{}) *RpcServer{
	return &RpcServer{Addr:addr,handle:handle,Status:true}
}

//建立 jsonrpc 服务
func (s *RpcServer)Run()  {
	rpc.Register(s.handle)
	address,e:=net.ResolveTCPAddr("tcp",":"+s.Addr)
	if e!=nil{
		log.Fatal(e)
	}
	//
	listener,err:=net.ListenTCP("tcp",address)
	if err!=nil{
		log.Fatal(e)
	}
	s.Status=true
	for  {
		if s.Status==false{
			fmt.Println("我退出了")
		}
		conn,err:= listener.Accept()
		if err!=nil{
			fmt.Println("小忙一下")
			continue
		}
		//服务起来了....

		jsonrpc.ServeConn(conn)
	}
}
