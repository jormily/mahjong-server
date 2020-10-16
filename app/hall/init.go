package db

import (
	"github.com/kudoochui/kudos/rpc"
	"github.com/kudoochui/kudos/service/msgService"
	"mahjong-server/protocol"
)

// register server service to remote
var msgArray = []interface{}{}

func RegisterHandler(msg interface{}){
	msgArray = append(msgArray, msg)
}

type MsgHandler struct {
	r rpc.HandlerRegister
}

func (m *MsgHandler)RegisterHandler()  {
	for _,v := range msgArray {
		m.r.RegisterHandler(v,"")
	}
}

func init() {
	RegisterHandler(new(Hall))

	//msgService.GetMsgService().Register("FindUser", &protocol.FindUserRequest{}, &protocol.FindUserResponse{})
	msgService.GetMsgService().Register("Hall.GetMessage", &protocol.MessageRequest{}, &protocol.MessageResponse{})
	msgService.GetMsgService().Register("Hall.GetUserStatus", &protocol.UserStatusRequest{}, &protocol.UserStatusResponse{})
	msgService.GetMsgService().Register("Hall.GetBaseInfo", &protocol.BaseInfoRequset{}, &protocol.BaseInfoResponse{})

}