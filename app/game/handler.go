package user

import (
	"context"
	"github.com/kudoochui/kudos/log"
	"github.com/kudoochui/kudos/rpc"
)

type HelloReq struct {
	Words	string
}

type HelloResp struct {
	Words	string
}

type Game struct {

}

func (h *Game) Say(ctx context.Context, args *rpc.Args, replay *HelloResp) error {
	var req HelloReq
	args.GetObject(&req)

	log.Info("hello" + req.Words)
	replay.Words = "hello " + req.Words

	return nil
}