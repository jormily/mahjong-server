package login

import (
	"context"
	"mahjong-server/util/crypto"

	rpcClient "github.com/kudoochui/kudos/component/proxy"
	//"github.com/kudoochui/kudos/log"
	"github.com/kudoochui/kudos/rpc"

	"mahjong-server/config"
	"mahjong-server/protocol"
)


type Login struct {

}


func (t *Login) ServerInfo(ctx context.Context, args *rpc.Args, reply *protocol.ServerInfoResponse) error {
	reply.Version,_ = config.GameConfig.Int("version")
	reply.Appweb = config.GameConfig.String("appWeb")

	return nil
}

func (this *Login) GuestAuth(ctx context.Context, args *rpc.Args, reply *protocol.GuestAuthResponse) error {
	var req protocol.GuestAuthRequest
	args.GetObject(&req)

	reply.Account = "guest_" + req.Account
	reply.Sign = string(crypto.MD5Digest([]byte(req.Account + args.Session.NodeAddr + config.GameConfig.String("accoutPriKey"))))
	return nil
}

func (this *Login) GuestLogin(ctx context.Context, args *rpc.Args, reply *protocol.LoginResponse) error {
	var req protocol.GuestAuthRequest
	args.GetObject(&req)

	proxy := GetLoginServer().GetComponent("proxy").(*rpcClient.Proxy)
	rpcArgs := &rpc.Args{
		Session: args.Session,
		MsgReq: protocol.FindUserRequest{
			Account: req.Account,
		},
	}
	rpcReplay := &protocol.FindUserResponse{}
	if err := proxy.RpcCall("DB","FindUser", rpcArgs, rpcReplay);err != nil {
		reply.Errcode = 1
		return err
	}

	reply.UserModel = rpcReplay.UserModel
	return nil
}