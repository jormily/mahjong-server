package user

import (
	"fmt"
	"github.com/kudoochui/kudos/app"
	rpcClient "github.com/kudoochui/kudos/component/proxy"
	rpcServer "github.com/kudoochui/kudos/component/remote"
	"github.com/kudoochui/kudos/component/timers"
	"github.com/kudoochui/kudos/log"
	"mahjong-server/config"
	"mahjong-server/manager/room"
	"mahjong-server/protocol"
	"time"
)

var (
	roomMgr = room.NewRoomManager()
)

type GameServer struct {
	*app.ServerDefault
	msgHandler *MsgHandler
}

var instance *GameServer

func GetGameServer() *GameServer {
	return instance
}

func init()  {
	app.RegisterCreateServerFunc("game", func(serverId string) app.Server {
		instance = &GameServer{
			ServerDefault: app.NewServerDefault(serverId),
		}
		return instance
	})
}

func (g *GameServer) Init(){
	settings, err := config.ServersConfig.GetMap("game")
	if err != nil {
		log.Error("%s", err)
	}
	serverSetting := settings[g.ServerId].(map[string]interface{})
	remoteAddr := fmt.Sprintf("%s:%.f",serverSetting["host"], serverSetting["port"])

	remote := rpcServer.NewRemote(
		rpcServer.Addr(remoteAddr),
		rpcServer.RegistryType(config.RegistryConfig.String("registry")),
		rpcServer.RegistryAddr(config.RegistryConfig.String("addr")),
		rpcServer.BasePath(config.RegistryConfig.String("basePath")))
	g.Components["remote"] = remote
	g.msgHandler = &MsgHandler{r:remote}


	proxy := rpcClient.NewProxy(
		rpcClient.RegistryType(config.RegistryConfig.String("registry")),
		rpcClient.RegistryAddr(config.RegistryConfig.String("addr")),
		rpcClient.BasePath(config.RegistryConfig.String("basePath")))
	g.Components["proxy"] = proxy
	//timer := timers.NewTimer()
	//log.Info("timer addr:%p",timer)
	g.Components["timer"] = timers.NewTimer()
	//timer = g.GetComponent("timer").(*timers.Timers)
	//log.Info("timer addr:%p",timer)

	g.ServerDefault.Init()
	// register service
	g.msgHandler.RegisterHandler()
}

func (g *GameServer)OnStart(){
	log.Info("OnStart")
	serverid,_ := config.SidsConfig.Int(g.ServerId)
	timer := g.GetComponent("timer").(*timers.Timers)
	proxy := g.GetComponent("proxy").(*rpcClient.Proxy)
	creater := func() bool {
		req := &protocol.FindRoomsRequest{ServerId: serverid}
		replay := &protocol.FindRoomsResponse{}
		if err := proxy.Call("DB.FindRooms",req,replay);err != nil {
			return false
		}
		//roomMgr.CreateRooms(serverid,replay.RoomList)

		return false
	}

	timer.AfterFuncEx(time.Second*5, creater)
}
