package rpc

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"os"
)

func TestClient(){
	client,err:=jsonrpc.Dial("tcp","127.0.0.1:1234")
	if err != nil {
		log.Fatal("Dial 发生了错误了哦 错误的信息为   err=",err)
	}
	var  resive  int
	type Argss struct {

		A, B int
	}
	args:=Argss{1,2}
	err1:=client.Call("Arith.Multiply",args,&resive)
	if err1!=nil {
		fmt.Println("shiming call error    ")
		fmt.Println("Call 的时候发生了错误了哦  err=",err1)
	}
	fmt.Println("收到信息了",resive)
	client.Close()
}

func Testsssss()  {
	service := "127.0.0.1:1234"
	client, err := jsonrpc.Dial("tcp", service)
	if err != nil {
		fmt.Println("dial error:", err)
		os.Exit(1)
	}
	type Argss struct {
		A, B int
	}
	args := Argss{1, 2}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		fmt.Println("Arith.Muliply call error:", err)
		os.Exit(1)
	}
	fmt.Println("the arith.mutiply is :", reply)
	client.Close()
}