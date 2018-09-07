package ipc

import (
	"fmt"
	"testing"
)
//定义一个输出服务端
type EchoServer struct {
}
//服务端当有客户端连接时的处理方法
func (server *EchoServer) Handel(method,params string) *Response{
	return &Response{"Ok","echo :"+method+"~"+params}
}
//返回服务器的名称
func (server *EchoServer) Name()string{
	return "server echo "
}

func TestIpc(t *testing.T)  {
	server:=NewIpcServer(&EchoServer{})
	clent1:=NewIpcClient(server)
	clent2:=NewIpcClient(server)
	resp1,_:=clent1.Call("foo","From client1")
	resp2,_:=clent2.Call("bar","from client2")
	fmt.Println(resp1)
	fmt.Println(resp2)
	if resp1.Body != "echo :foo~From client1"||resp2.Body!="echo :bar~from client2"{
		t.Error("ipcClient.call faild")
	}
	//defer clent2.Close()
	//defer clent1.Close()
}



