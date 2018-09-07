package rpc

import (
	"fmt"
	"testing"
)

type Argss struct {
	A,B int
}

type Quotient struct {
	Quo,Rem int
}

type Arith int

func (a *Arith) Multiply(args *Argss,reply *int ) error  {
	*reply=args.A*args.B
	return nil
}

var defaultServer *RpcServer

func TestRpcServer(t *testing.T) {
	a:=new(Arith)
	serverManaer:=NewServerManager() // 创建一个服务器管理器
	defaultServer=NewRpcServer("1234",a)
	serverManaer.Register("default",defaultServer)

	//serverManaer.Register("test",)//将某个对象作为 rpc服务去注册

	serverManaer.Run("") //跑起来...
	fmt.Println(serverManaer.Close("test"))
	//t.Run(a)
}