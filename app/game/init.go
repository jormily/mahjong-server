package user

import (
	"github.com/kudoochui/kudos/rpc"
	_ "mahjong-server/protocol"
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
	//RegisterHandler(new(Hel))
	//
	//// register msg type
	//msgService.GetMsgService().Register("Hel.Say", &protocol.EmptyRequest{}, &protocol.EmptyResponse{})

}
