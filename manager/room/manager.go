package room

import (
	"fmt"
	"github.com/kudoochui/kudos/rpc"
	"mahjong-server/protocol"
	"mahjong-server/util"
	"strconv"
)

type RoomManager struct {
	location 	map[int64]*Location
	rooms		map[string]*Room
	creatingRooms map[string]bool
	totalRooms	int
}

func NewRoomManager() *RoomManager {
	r := new(RoomManager)
	r.location = map[int64]*Location{}
	r.rooms = map[string]*Room{}
	r.creatingRooms = map[string]bool{}
	r.totalRooms = 0
	return r
}

func checkConfig(conf *protocol.RoomCnf) bool {
	return true
}

func (r *RoomManager) GenerateRoomId(check func(id string) bool) string {
	var roomid string

	generateId := func() string {
		roomId := 0
		for i := 0;i < 6;i++ {
		roomId = roomId*10 + (util.Rand(1,9))
	}
		return strconv.FormatInt(int64(roomId),10)
	}

	for {
		roomid = generateId()
		if check(roomid) {
			continue
		}else{
			break
		}
	}
	return roomid
}

func (r *RoomManager) CreateRoom(rm *protocol.RoomModel) {
	if _,ok := r.rooms[rm.Id];ok {
		return
	}

	rr := NewRoom(rm)
	r.rooms[rm.Id] = rr
}

func (r *RoomManager) CreateRooms(serverid int,rms []protocol.RoomModel){
	for _,rm := range rms {
		r.CreateRoom(&rm)
	}
}

func (r *RoomManager) GetUserLocation(userid int64) (*Location,bool) {
	if l,ok := r.location[userid];ok {
		return l,true
	}
	return nil,false
}

func (r *RoomManager) GetRoomByUserId(userid int64) *Room {
	if l,ok := r.GetUserLocation(userid);ok {
		if r,ok := r.rooms[l.roomid];ok {
			return r
		}
	}
	return nil
}

func (r *RoomManager) GetRoom(roomid string) *Room {
	if r,ok := r.rooms[roomid];ok {
		return r
	}
	return nil
}


func (r *RoomManager) EnterRoom(s *rpc.Session,roomid string,userid int64,name string,score int) error {
	l,ok := r.GetUserLocation(userid)
	if ok && l.roomid == roomid {
		return fmt.Errorf("room exist error")
	}

	if rm,ok := r.rooms[roomid];ok {
		if index := rm.PlayerEnter(s,name,score);index >= 0 {
			r.location[userid] = &Location{
				roomid: roomid,
				seate: index,
			}
			return nil
		}else{
			return fmt.Errorf("player enter error")
		}
	}else{
		return fmt.Errorf("room status error")
	}
}

