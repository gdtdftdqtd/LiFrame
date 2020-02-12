package main

import (
	"LiFrame/core/liNet"
	"LiFrame/proto"
	"LiFrame/server/app"
	"LiFrame/server/db"
	"LiFrame/server/world"
	"LiFrame/utils"
	"os"
)


func main() {

	if len(os.Args) > 1 {
		cfgPath := os.Args[1]
		utils.GlobalObject.Load(cfgPath)
	}else{
		utils.GlobalObject.Load("conf/world.json")
	}

	db.InitDataBase()

	s := liNet.NewServer()
	s.AddRouter(&world.Enter)
	s.AddRouter(&world.CommonWorld)

	s.SetOnConnStart(world.ClientConnStart)
	s.SetOnConnStop(world.ClientConnStop)
	app.SetServer(s)
	app.SetShutDownFunc(world.ShutDown)

	go app.MasterClient(proto.ServerTypeWorld)

	s.Running()
}
