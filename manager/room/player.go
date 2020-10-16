package room

import (
	"github.com/kudoochui/kudos/rpc"
)

type Player struct {
	// 基础数据
	Session 	*rpc.Session
	UserId 		int64
	Score 		int
	Name 		string

	// 状态数据
	State 		int
	Ready		bool
	Seate		int

	// 游戏数据
	NumZiMo		int 	//自摸
	NumJiePao	int
	NumDianPao	int
	NumAnGang	int
	NumMingGang	int
	NumChaJiao	int
}

func (p *Player) init() {
	p.Session = nil
	p.UserId = 0
	p.Score = 0
	p.Name = ""

	p.State = 0
	p.Ready = false
}