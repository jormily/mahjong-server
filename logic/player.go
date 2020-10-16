package logic

import (
	"github.com/kudoochui/kudos/rpc"
	. "mahjong-server/util/slice"
)

type Player struct {
	*rpc.Session
	seat 		int

	holds 		[]int
	folds 		[]int
	agCards 	[]int
	dgCards 	[]int
	wgCards 	[]int
	peCards		[]int
	lack		int
	changeCards []int
	gangCards 	[]int
	huCard 		int
	huList 		[]int

	cardMap 	map[int]int
	tingMap 	map[int]int

	canGang		bool
	canPen		bool
	canHu		bool
	isZiMo		bool
	canPlay		bool
	state 		int


	hufan	 	int
	drawed		bool
	drawS 		bool
	drawG		bool

	action 		[]int
	fan 		int
	score  		int

	lastfg 		int		// 上一次放杠玩家

	numZiMo 	int
	numJiePao 	int
	numDianPao 	int
	numAnGang 	int
	numMingGang int
	numChaJiao 	int
}

func newPlayer(seat int) *Player {
	this := new(Player)
	this.seat = seat
	//this.init()
	return this
}

func (this *Player) setState(idx int,state int){
	if state == 0 {
		this.state = this.state & (1<<idx)
	}
}

func (this *Player) getHoldCardType() [3]int {
	tl := [3]int{}
	for _,card := range this.holds {
		ct := getCardType(card)
		tl[ct]++
	}
	return tl
}

func (this *Player) checkCards(c interface{}) bool {
	switch c.(type) {
	case int:
		if this.cardMap[c.(int)] > 0 {
			return true
		}else{
			return false
		}
	case []int:
		cardMap := Slice2Map(c.([]int))
		for card,cnt := range cardMap {
			if this.cardMap[card]  < cnt {
				return false
			}
		}
		return true
	default:
		return false

	}
}

func (this *Player) checkCardLack(card int) bool {
	return this.lack == getCardType(card)
}

/**
是否打缺
*/
func (this *Player) checkLack() bool {
	for _,card := range this.holds {
		if getCardType(card) == this.lack {
			return false
		}
	}
	return true
}

func (this *Player) checkTing() bool {
	if !this.checkLack() {
		return false
	}

	return true
}

func (this *Player)init(){
	this.holds 	= []int{}
	this.folds 	= []int{}
	this.agCards = []int{}
	this.dgCards = []int{}
	this.wgCards = []int{}
	this.peCards	= []int{}
	this.huList = []int{}

	this.lack = -1
	this.hufan = -1

	this.changeCards = nil
	this.action =	[]int{}
	this.cardMap =	map[int]int{}
	this.tingMap =	map[int]int{}

	this.canGang =	false
	this.canPen =	false
	this.canHu =	false
	this.canPlay =	false
	this.drawed =	false
	this.drawS = 	false
	this.drawG =	false
}
