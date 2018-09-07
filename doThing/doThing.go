package doThing

import "fmt"

type DoThing struct {
}

func (server *DoThing) SayHello() {
	fmt.Println("Server FunOne")
}
