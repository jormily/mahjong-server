package db

import (
	"context"
	rpcClient "github.com/kudoochui/kudos/component/proxy"
	"github.com/kudoochui/kudos/log"
	"github.com/kudoochui/kudos/rpc"
	"mahjong-server/protocol"
)

type Hall struct {

}

func (h *Hall) GetMessage(ctx context.Context, args *rpc.Args, replay *protocol.MessageResponse) error {
	log.Info("GetMessage")
	var req protocol.MessageRequest
	args.GetObject(&req)

	proxy := GetHallServer().GetComponent("proxy").(*rpcClient.Proxy)
	rpcArgs := &rpc.Args{
		Session: args.Session,
		MsgReq: protocol.FindMessageRequset{
			Type: req.Type,
			Version: req.Version,
		},
	}

	rpcReplay := &protocol.FindMessageResponse{}
	if err := proxy.RpcCall("DB","FindMessage", rpcArgs, rpcReplay);err != nil {
		replay.Errcode = 1
		return err
	}

	replay.MessageModel = rpcReplay.MessageModel
	log.Info("%v",replay)
	return nil
}

func (h *Hall) GetUserStatus(ctx context.Context, args *rpc.Args, replay *protocol.UserStatusResponse) error {
	log.Info("GetUserStatus")
	var req protocol.UserStatusRequest
	args.GetObject(&req)

	proxy := GetHallServer().GetComponent("proxy").(*rpcClient.Proxy)
	rpcArgs := &rpc.Args{
		Session: args.Session,
		MsgReq: protocol.FindUserRequest{
			Account: req.Account,
		},
	}
	rpcReplay := &protocol.FindUserResponse{}
	if err := proxy.RpcCall("DB","FindUser", rpcArgs, rpcReplay);err != nil {
		replay.Errcode = 1
		return err
	}

	replay.Gems = int(rpcReplay.Gems)
	log.Info("%v",replay)
	return nil
}

func (h *Hall) GetBaseInfo(ctx context.Context, args *rpc.Args, replay *protocol.BaseInfoResponse) error {
	log.Info("GetBaseInfo")
	var req protocol.BaseInfoRequset
	args.GetObject(&req)

	proxy := GetHallServer().GetComponent("proxy").(*rpcClient.Proxy)
	rpcArgs := &rpc.Args{
		Session: args.Session,
		MsgReq: protocol.FindUserRequest{
			UserId: req.UserId,
		},
	}
	rpcReplay := &protocol.FindUserResponse{}
	if err := proxy.RpcCall("DB","FindUser", rpcArgs, rpcReplay);err != nil {
		replay.Errcode = 1
		return err
	}

	replay.Sex = int(rpcReplay.Sex)
	replay.Name = rpcReplay.Name
	replay.HeadImgUrl = rpcReplay.Headimg
	log.Info("%v",replay)
	return nil
}