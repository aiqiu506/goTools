package service

import (
	"fmt"
	"reflect"
)

// ReServer 来保存map的结构体
type ReServer struct {
	M map[string]reflect.Value
}

// RegisterService 注册服务
func (this *ReServer) RegisterService(service interface{}) (err error) {
	serviceType := reflect.TypeOf(service)
	serviceValue := reflect.ValueOf(service)
	ServiceNum := serviceType.NumMethod()
	for i := 0; i < ServiceNum; i++ {
		ServiceName := serviceType.Method(i).Name
		if _, ok := this.M[ServiceName]; ok {
			fmt.Println("service has been registered")
		} else {
			this.M[ServiceName] = serviceValue.Method(i)
		}
	}

	return
}

//  Start 服务启动
func (this *ReServer) Start(m string) {
	if _,ok:=this.M[m];ok{
		this.M[m].Call(nil)
	}else{
		fmt.Println("方法不存在!")
	}

}
