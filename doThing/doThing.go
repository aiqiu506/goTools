package doThing

import "fmt"

type DoThing struct {
}

func (server *DoThing) FunOne(){
	fmt.Println("Server FunOne")
}
