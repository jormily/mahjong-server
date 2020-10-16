package db

import (
	"github.com/kudoochui/kudos/rpc"
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
	RegisterHandler(new(DB))

	//msgService.GetMsgService().Register("Hall.GetMessage", &protocol.MessageRequest{}, &protocol.MessageResponse{})
	//msgService.GetMsgService().Register("Hall.GetUserStatus", &protocol.UserStatusRequest{}, &protocol.UserStatusResponse{})

}