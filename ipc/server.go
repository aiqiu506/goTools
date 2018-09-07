package ipc

import (
	"encoding/json"
	"fmt"
)

// 请求对象
type Request struct {
	 Method string `json:"method"`
	 Params string `json:"params"`
}
//返回对象
type Response struct {
	 Code string `json:"code"`
	 Body string `json:"body"`
}
//服务器接口
 type Server interface {
 	 Name() string
 	 Handel(method,params string) *Response
 }
//ipc服务器对象

type IpcServer struct {
	Server
}

//新建一个 ipc服务器对象
func NewIpcServer(s Server) *IpcServer{
	return &IpcServer{s}
}

//建立一个 ipc服务器的连接
func (server *IpcServer) Connect() chan string{
	//创建一个 channel 通信通道
	session:=make(chan string,0)
	go func(c chan string) {
		for  {
			//从通信通道中读出数据
			request:=<-c
			if request=="CLOSE"{
				break
			}
			var req Request
			//将读出的数据 反序列化
			err:=json.Unmarshal([]byte(request),&req)
			if err!=nil{
				fmt.Println("Invalid request format:",err)
			}
			//调用服务器实现的 handel 方法处理请求,并返回一个 Response 的对象
			resp:=server.Handel(req.Method,req.Params)
			//将返回的 Response 对象再序列化后,进行响应输出
			b,err:=json.Marshal(resp)
			//输出到 session 中,返回给客户端
			c <- string(b)
		}
		fmt.Println("session closed...")
	}(session)
	fmt.Println("A new session has been created successfully")
	return session
}
