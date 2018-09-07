package cg

import (
	"fmt"
)

type Player struct {
	Name string
	Level int
	Exp int
	Room int
	mq chan * Message
}

func NewPlayer() *Player{
	m:=make(chan *Message,1024)
	p:=&Player{"",0,0,0,m}
	go func(player *Player) {
		for{
			msg:=<-player.mq
			fmt.Println(player.Name,"received message:"+msg.Content)
		}
	}(p)
	return p
}


