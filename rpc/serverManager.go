package rpc

import "fmt"

//服务器管理

type ServerManager struct {
	servers map[string]*RpcServer
}

//新建服务器
func NewServerManager()*ServerManager{
	return &ServerManager{make(map[string]*RpcServer)}
}


//将服务器注册进来
func (serverM *ServerManager)Register(name string,s *RpcServer){
	if _,isExist:=serverM.servers[name];isExist{
		return
	}
	serverM.servers[name]=s
}

func (serverM *ServerManager)Run(name interface{}){
	if name!=nil{
		//启动所有注册过的服务器
		for serverName,server:= range serverM.servers{
			if server.Status==false{
				continue
			}
			fmt.Println("服务器:"+serverName,",已经启动")
			server.Run()
		}
	}else{
		//启动单个服务器
		if server,isok:=serverM.servers[name.(string)];isok&&serverM.servers[name.(string)].Status==true{
			fmt.Println("服务器:"+name.(string),",已经启动")
			server.Run()
		}

	}
}

//关闭某一服务器
func (serverM *ServerManager)Close(name string)bool{
	if _,isExist:=serverM.servers[name];isExist{
		return false
	}
	serverM.servers[name].Status=false
	return true
}
