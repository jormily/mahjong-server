package session

import (
	"github.com/kudoochui/kudos/rpc"
)

type SessionManager struct {
	smap  	map[int64]*rpc.Session
}

func NewSessionManager() *SessionManager {
	sm := new(SessionManager)
	sm.smap = make(map[int64]*rpc.Session)
	return sm
}

func (sm *SessionManager)Add(s *rpc.Session,userid int64){
	s.Bind(userid)
	sm.smap[userid] = s
}

func (sm *SessionManager)Delete(userid int64){
	if s,ok := sm.smap[userid];ok {
		s.UnBind()
		delete(sm.smap,userid)
	}
}