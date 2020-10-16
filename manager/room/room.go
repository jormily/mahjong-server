package room

import (
	"encoding/json"
	"github.com/kudoochui/kudos/rpc"
	"github.com/kudoochui/kudos/service/channelService"
	"mahjong-server/logic"
	"mahjong-server/protocol"
	"math/rand"
	"strconv"
	"time"
)

func generateRoomId() string {
	roomId := 0
	for i := 0;i < 6;i++ {
		rand.Seed(time.Now().UnixNano())
		roomId = roomId*10 + (rand.Intn(9) + 1)
	}
	return strconv.FormatInt(int64(roomId),10)
}

type Room struct {
	channel 		*channelService.Channel
	uuid 			string
	id 				string
	turns			int
	createTime		int64
	nextButton 		int
	players			[]*Player
	conf 			protocol.RoomCnf
	logic 			*logic.Manager
}

func NewRoom(rm *protocol.RoomModel) *Room {
	r := new(Room)
	r.uuid = rm.Uuid
	r.id = rm.Id
	r.turns = 0
	r.createTime = rm.CreateTime
	r.nextButton = int(rm.NextButton)
	r.players =  make([]*Player,4)
	r.channel = channelService.NewChannel(rm.Id)
	r.logic = logic.NewManager(r)

	json.Unmarshal([]byte(rm.Config),&r.conf)

	for i:=0;i<len(r.players);i++ {
		r.players[i] = new(Player)
		r.players[i].State = 0
		r.players[i].Seate = i
	}

	return r
}

/*
get player userid
 */
func (r *Room)GetUserId(seat int) int64 {
	if r.players[seat] == nil {
		return 0
	}
	return r.players[seat].UserId
}

/*
get player id
 */
func (r *Room)GetPlayerId(userid int64) int{
	for _,v := range r.players {
		if v.UserId == userid {
			return v.Seate
		}
	}
	return  -1
}

/**
set all player ready false
 */
func (r *Room)ClearReady(){
	for _,v := range r.players {
		v.Ready = false
	}
}

/*
get player
 */
func (r *Room)GetPlayer(id interface{}) *Player {
	switch id.(type) {
	case int:
		if id.(int) < 0 || id.(int) >= 4{
			return nil
		}
		return r.players[id.(int)]
	case int64:
		for _,v := range r.players {
			if v.UserId == id.(int64) {
				return v
			}
		}
		return nil
	default:
		return nil
	}
}

/**
set player ready
 */
func (r *Room)SetReady(id int,ready bool){
	p := r.GetPlayer(id)
	p.Ready = ready
}

/**
check ready
 */
func (r *Room)CheckReady() bool {
	for _,p := range r.players {
		if !p.Ready {
			return false
		}
	}
	return true
}

func (r *Room)GetAllPlayer() []*Player {
	return r.players[:]
}

/**
logic model
 */
func (r *Room)GetLogic() *logic.Manager {
	return r.logic
}

/**
player enter
 */
func (r *Room)PlayerEnter(s *rpc.Session,name string,score int) int {
	for _,v := range r.players {
		if v.UserId == 0 {
			v.UserId = s.GetUserId()
			v.Name = name
			v.Score = score
			v.Session = s
			r.channel.Add(s)
			return v.Seate
		}
	}
	return -1
}

/**
player leave
 */
func (r *Room)PlayerLeave(id interface{}) bool {
	switch id.(type) {
	case int:
		if id.(int) >= 0 && id.(int) < len(r.players) {
			for _,v := range r.players {
				if v.Seate == id.(int) {
					r.channel.Leave(v.UserId)
					v.init()
					return true
				}
			}
		}
		return false
	case int64:
		for _,v := range r.players {
			if v.UserId == id.(int64) {
				r.channel.Leave(v.UserId)
				v.init()
				return true
			}
		}
		return false
	default:
		return false
	}
}

/**
send message
 */
func (r *Room)PushMessage(index int,route string,msg interface{}){
	if index < 0 || index >= len(r.players) {
		return
	}
	p := r.players[index]
	p.Session.PushMessage(route,msg)
}

/**
broadcast message
*/
func (r *Room)Broadcast(route string,msg interface{}){
	r.channel.PushMessage(route,msg)
}

/**
broadcast message except id
*/
func (r *Room)BroadcastEx(route string,msg interface{},id int){
	for _,p := range r.players {
		if  p.Seate != id {
			p.Session.PushMessage(route,msg)
		}
	}
}



