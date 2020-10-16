package login

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
	RegisterHandler(new(Login))

	msgService.GetMsgService().Register("Login.ServerInfo", &protocol.EmptyRequest{}, &protocol.ServerInfoResponse{})
	msgService.GetMsgService().Register("Login.GuestAuth", &protocol.GuestAuthRequest{}, &protocol.GuestAuthResponse{})
	msgService.GetMsgService().Register("Login.GuestLogin", &protocol.LoginRequest{}, &protocol.LoginResponse{})
}