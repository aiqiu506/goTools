package ipc

import "encoding/json"

//定义客户端对象  客户端对象,conn字段为一个 channel,用于保存从服务端连接后的通道
 type IpcClient struct {
 	conn chan string
 }

//建立一个 ipc 客户端的连接 从 参数 server 服务端的 connect 方法里,将通信通道获取到,并将该通道用于生成一个新的客户端连接对象
func NewIpcClient(server *IpcServer) *IpcClient{
    c:=server.Connect()  //返回一个通信通道
	return &IpcClient{c}
}

//客户端向服务端通信,发送请求,处理响应,   参数为 method,params ,用其构建一个 请求对象, 并将这个对象序列化成json 作为通信的数据
func (client * IpcClient) Call(method,params string) (resp *Response,err error){
	req:=&Request{method,params}
	var b []byte
	b,err=json.Marshal(req)
	if err!=nil{
		return
	}
	//将请求的数据,发送或写入 客户端的通信通道里
	client.conn<-string(b)
	//将通道返回的数据读出 再序列化返回
	str:=<-client.conn
	var resp1 Response
	err=json.Unmarshal([]byte(str),&resp1)
	resp=&resp1
	return
}
//发送关闭信号
func (client *IpcClient) Close(){
	client.conn<-"CLOSE"
}


