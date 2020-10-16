package db


import (
	"fmt"
	"github.com/kudoochui/kudos/app"
	rpcClient "github.com/kudoochui/kudos/component/proxy"
	rpcServer "github.com/kudoochui/kudos/component/remote"
	"github.com/kudoochui/kudos/log"
	"mahjong-server/app/db/database"
	"mahjong-server/config"
)

type DBCloser func()

type DBServer struct {
	*app.ServerDefault

	msgHandler 	*MsgHandler
	dbCloser	DBCloser
}


var instance *DBServer

func GetGateServer() *DBServer {
	return instance
}

func init()  {
	app.RegisterCreateServerFunc("db", func(serverId string) app.Server {
		instance = &DBServer{
			ServerDefault: app.NewServerDefault(serverId),
		}
		return instance
	})
}

func (g *DBServer) Init(){
	settings, err := config.ServersConfig.GetMap("db")
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

	maxIdle,_ := config.MysqlConfig.Int("maxIdle")
	maxOpen,_ := config.MysqlConfig.Int("maxOpen")
	showSql,_ := config.MysqlConfig.Bool("showSql")
	g.dbCloser = database.DBStartup(
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
			config.MysqlConfig.String("user"),
			config.MysqlConfig.String("password"),
			config.MysqlConfig.String("host"),
			config.MysqlConfig.String("port"),
			config.MysqlConfig.String("database"),
			config.MysqlConfig.String("args")),
		maxIdle,
		maxOpen,
		showSql)
	// register service
	g.ServerDefault.Init()

	g.msgHandler.RegisterHandler()
}

func (g *DBServer) OnStop(){
	for _,com := range g.Components {
		com.OnDestroy()
	}

	g.dbCloser()
}