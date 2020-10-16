package gate

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
	// register service
	//RegisterHandler(new(Gate))
	//
	//// register msg type
	//msgService.GetMsgService().Register("Gate.ServerInfo", &Args{}, &Reply{})
	//msgService.GetMsgService().RegisterNotify("Arith.TestRpc", &HelloReq{})
}