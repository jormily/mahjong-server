package gate

import (
	"context"
	"fmt"
	rpcClient "github.com/kudoochui/kudos/component/proxy"
	"github.com/kudoochui/kudos/log"
	"github.com/kudoochui/kudos/rpc"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type HelloReq struct {
	Words	string
}

type HelloResp struct {
	Words	string
}

type String string


type Gate struct {

}


func (t *Gate) Mul(ctx context.Context, args *rpc.Args, reply *Reply) error {
	var msgReq Args
	err := args.GetObject(&msgReq)
	if err != nil {
		log.Error("%+v", err)
	}

	reply.C = msgReq.A * msgReq.B
	fmt.Printf("call: %d * %d = %d\n", msgReq.A, msgReq.B, reply.C)
	return nil
}

func (t *Gate) Add(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	fmt.Printf("call: %d + %d = %d\n", args.A, args.B, reply.C)
	return nil
}

func (t *Gate) TestRpc(ctx context.Context, args *rpc.Args,reply *Reply) error {
	var req HelloReq
	args.GetObject(&req)
	log.Info("RPCTest " + req.Words)
	//reply.Words = "hello " + req.Words
	proxy := GetGateServer().GetComponent("proxy").(*rpcClient.Proxy)

	rep := &HelloResp{}
	if err := proxy.RpcCall("Hi", "TestRpc", args, rep);err != nil {
		log.Info(err.Error())
	}

	log.Info("%v",rep)

	return nil
}