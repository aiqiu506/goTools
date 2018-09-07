package grqStruct

import (
	"testing"
	"fmt"
)

func TestStructToMap(t *testing.T) {
	type s struct {
		A string
		C int
	}
	type A struct {
		name string
		Id int
		V  s
		Sli []string
	}

	a:=A{
			name:"abc",
			Id:1,
			V:s{ "abc", 1},
			Sli:[]string{"A","B","C"},
		}


	data:=StructToMap(a)
	fmt.Printf("%v",data)

}