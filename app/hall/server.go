package db

import (
	"fmt"
	"github.com/kudoochui/kudos/app"
	rpcClient "github.com/kudoochui/kudos/component/proxy"
	rpcServer "github.com/kudoochui/kudos/component/remote"
	"github.com/kudoochui/kudos/log"
	"mahjong-server/config"
)

type HallServer struct {
	*app.ServerDefault

	msgHandler 	*MsgHandler
}

var instance *HallServer

func GetHallServer() *HallServer {
	return instance
}

func init()  {
	app.RegisterCreateServerFunc("hall", func(serverId string) app.Server {
		instance = &HallServer{
			ServerDefault: app.NewServerDefault(serverId),
		}
		return instance
	})
}

func (g *HallServer) Init(){
	settings, err := config.ServersConfig.GetMap("hall")
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

	g.ServerDefault.Init()
	// register service
	g.msgHandler.RegisterHandler()
}

