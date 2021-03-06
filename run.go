package main

import (
	"fmt"
	"goTools/doThing"
	"goTools/service"
	"reflect"
)

func main() {
	//引用service包中的ReServer类型来创建reServer变量，目的是使reServer 变量可以调用RegisterServer方法来将该包的services类型的变量进行注入
	reServer := &service.ReServer{
		M: make(map[string]reflect.Value),
	}
	var doThings doThing.DoThing
	err := reServer.RegisterService(&doThings)
	if err!=nil{
		fmt.Println(err.Error())
	}

	reServer.Start("SayHello")

	/*
	 测试  两浮点数比较
	*/
	//a,b:=0.1234,0.12340
	//fmt.Println(tool.IsEqual(a,b,0.0000001))

	/**
	测试 取子串
	*/
	//str:="abc你好db"
	//sub,err:=tool.SubStr(str,-3,1)
	//fmt.Println([]rune(str),sub,err)

	/*
		测试 判断是否在数组中
	*/
	//1.简单元素判断
	//arr1 := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(tool.InArray(arr1, "1"))
	//fmt.Println(tool.InArray(arr1,k))
	//2 结构体元素判断
	//type s struct {
	//	Id int
	//	Name string
	//}
	//arr:=[]s{{Id:1,Name:"a"},{Id:3,Name:"b"},{Id:4,Name:"c"} }
	//p:=s{Id:3,Name:"b"}
	//fmt.Println(tool.InArray(arr,p))

	/**
	测试格式化时间
	*/
	//fmt.Println(tool.FormaterTime("2006-01-02 15:04"))
}
