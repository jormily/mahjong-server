package db

import (
	"context"
	"fmt"
	"github.com/kudoochui/kudos/log"
	"github.com/kudoochui/kudos/rpc"
	"mahjong-server/app/db/database"
	"mahjong-server/protocol"
)

type DB struct {

}

func (d *DB) FindUser(ctx context.Context, args *rpc.Args, replay *protocol.FindUserResponse) error {
	var req protocol.FindUserRequest
	args.GetObject(&req)

	var user *database.TUsers
	var err error
	if req.UserId != 0 {
		user,err = database.FindUserByUserId(req.UserId)
		if err != nil {
			return err
		}
	}else if req.Account != "" {
		user,err = database.FindUserByUserId(req.UserId)
		if err != nil {
			return err
		}
	}else {
		return fmt.Errorf("requset data is error")
	}

	replay.UserId = user.Userid
	replay.Account = user.Account
	replay.Name = user.Name
	replay.Level = user.Lv
	replay.Exp = user.Exp
	replay.Coins = user.Coins
	replay.Gems = user.Gems
	replay.Sex = user.Sex
	replay.RoomId = user.Roomid
	replay.Headimg = user.Headimg

	return nil
}


func (d *DB) FindMessage(ctx context.Context, args *rpc.Args, replay *protocol.FindMessageResponse) error {
	var req protocol.FindMessageRequset
	args.GetObject(&req)

	message,err := database.FindMessage(req.Type,req.Version)
	if err != nil {
		return err
	}

	replay.MessageModel.Version = message.Version
	replay.MessageModel.Type = message.Type
	replay.MessageModel.Msg = message.Msg

	return nil
}

/*
server call
 */
func (d *DB) FindRooms(ctx context.Context,args *rpc.Args, replay *protocol.FindRoomsResponse) error {
	log.Info("FindRooms")
	var req protocol.FindRoomsRequest
	args.GetObject(&req)

	replay.RoomList = make([]protocol.RoomModel,0)
	for _,rm := range database.FindRoomByServerId(int64(req.ServerId)) {
		r := protocol.RoomModel{}
		r.Id = rm.Id
		r.Uuid = rm.Uuid
		r.Config = rm.BaseInfo
		r.CreateTime = rm.CreateTime
		r.NextButton = rm.NextButton
		replay.RoomList = append(replay.RoomList,r)
	}
	return nil
}