package cg

import (
	"goTools/ipc"
	"sync"
)

type CenterServer struct {
	servers map[string] ipc.Server
	players map[string]*Player
	rooms []*Room
	locker sync.RWMutex
}

type Room struct {
	ID int
}


type Message struct {
	From string `json:"from"`
	To string   `json:"to"`
	Content string `json:"content"`
}

func NewCenterServer () *CenterServer{
	servers:=make(map[string] ipc.Server)
	players:=make(map[string] *Player)
	rooms:=make([]*Room,0)
	return &CenterServer{servers:servers,players:players,rooms:rooms}
}

//添加玩家
func (uc * CenterServer) addPlayer(name string) error{
	p:=NewPlayer()
	uc.locker.Lock()
	defer uc.locker.Unlock()
	p.Name=name
	uc.players[name]=p
	return nil
}
func (uc *CenterServer) delPlayer(name string){
    delete (uc.players,name)
}

