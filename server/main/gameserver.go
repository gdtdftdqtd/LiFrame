package main

import (
	"LiFrame/core/liNet"
	"LiFrame/proto"
	"LiFrame/server/app"
	"LiFrame/server/db"
	"LiFrame/server/game"
	"LiFrame/utils"
	"os"
)


func main() {

	if len(os.Args) > 1 {
		cfgPath := os.Args[1]
		utils.GlobalObject.Load(cfgPath)
	}else{
		utils.GlobalObject.Load("conf/game.json")
	}

	db.InitDataBase()

	s := liNet.NewServer()
	s.AddRouter(&game.Enter)

	s.SetOnConnStart(game.ClientConnStart)
	s.SetOnConnStop(game.ClientConnStop)
	app.SetShutDownFunc(game.ShutDown)
	app.SetServer(s)

	go app.MasterClient(proto.ServerTypeGame)

	s.Running()
}
