package logic

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kudoochui/kudos/log"
	"github.com/kudoochui/kudos/rpc"
	"mahjong-server/protocol"
	. "mahjong-server/util"
	. "mahjong-server/util/intlist"
	. "mahjong-server/util/slice"
)



type Manager struct {
	round 			int
	cardList 		[]int
	cardIndex 		int
	banker 			int
	turn 			int
	state 			string
	actionList 		[]int
	drawCardList 	[]int
	playerList 		[]*Player
	config 			protocol.RoomCnf

	playCnt 		int
	playCard 		int

	r			Room
}

/**
new logic class
 */
func NewManager(r Room) *Manager {
	m := new(Manager)
	m.round = 0
	m.cardList = make([]int,0)
	m.banker = -1
	m.turn = -1
	m.playerList = make([]*Player,4)
	m.r = r
	m.init()

	return m
}

func (m *Manager) GetPlayer(id int) *Player {
	return m.playerList[id]
}

func (m *Manager)GetPlayerUserId(id int) int64 {
	player := m.GetPlayer(id)
	return player.Session.GetUserId()
}

func (m *Manager)PlayerEnter(s *rpc.Session,id int){
	p := m.GetPlayer(id)
	p.Session = s
}

func (m *Manager)PlayerLeave(id int){
	p := m.GetPlayer(id)
	p.Session = nil
	p.init()
}


/**
init logic
 */
func (m *Manager)init(){
	m.drawCardList = []int{}
	m.state = "idle"
	m.playCnt = 0
	m.playCard = -1
	m.round ++

	if m.banker == -1 {
		m.banker = Rand(0,3)
	}else{
		m.banker = (m.banker+1)%4
	}
	m.turn = m.banker

	for i := 0;i < 4; i ++ {
		if m.playerList[i] == nil {
			m.playerList[i] = newPlayer(i)
		}
		m.playerList[i].init()
	}

	if len(m.cardList) == 0 {
		for i:=Card_Value_Start;i<=Card_Value_End;i++ {
			for j:=0;j<4;j++ {
				m.cardList = append(m.cardList,i)
			}
		}
	}
	m.cardIndex = 0
}


func (m *Manager) checkPeng(id int,card int) {
	for k,v := range m.playerList {
		if k == id || v.drawed {
			continue
		}

		if v.checkCardLack(card) {
			continue
		}

		if v.cardMap[card] >= 2 {
			v.canPen = true
		}
	}

}

func (m *Manager) checkDianGang(id int,card int) {
	//如果没有牌了，则不能再杠
	if len(m.cardList) <= m.cardIndex {
		return
	}

	for k,v := range m.playerList {
		if v.drawed || k == id || v.checkCardLack(card) {
			continue
		}

		if v.cardMap[card] >= 3 {
			v.gangCards = append(v.gangCards, card)
			v.canGang = true
		}
	}
}

func (m *Manager) checkHu(id int,card int) {
	if card >= 0 {
		for sid,player := range m.playerList {
			if sid != id {
				//log.Info("id = %d",sid)
				if player.drawed {
					continue
				}

				cards := append([]int{}, player.holds...)
				cards = append(cards, card)

				if checkHu(cards,player.lack) {
					player.canHu = true
					player.huCard = card
					player.isZiMo = false
				}
			}
		}
	}else {
		//log.Info("id = %d",id)
		player := m.GetPlayer(id)
		if checkHu(player.holds,player.lack) {
			player.canHu = true
			player.huCard = player.holds[len(player.holds)-1]
			player.isZiMo = true
		}
	}
}

func (m *Manager) checkAnGang(id int) {
	if len(m.cardList) <= m.cardIndex {
		return
	}
	player := m.GetPlayer(id)
	for card,cnt := range player.cardMap {
		if !player.checkCardLack(card) && cnt == 4 {
			player.gangCards = append(player.gangCards,card)
			player.canGang = true
		}
	}
}

func (m *Manager) checkWanGang(id int) {
	if len(m.cardList) <= m.cardIndex {
		return
	}
	player := m.GetPlayer(id)
	for _,card := range player.peCards {
		if player.cardMap[card] > 0 {
			player.gangCards = append(player.gangCards,card)
			player.canGang = true
		}
	}
}

func (m *Manager) clearOpt(id int) {
	if id >= 0 && id < len(m.playerList) {
		player := m.GetPlayer(id)
		player.canHu = false
		player.canGang = false
		player.canPen = false
		if len(player.gangCards) > 0 {
			player.gangCards = []int{}
		}
	}else{
		for _, v := range m.playerList {
			v.canHu = false
			v.canGang = false
			v.canPen = false
			if len(v.gangCards) > 0 {
				v.gangCards = []int{}
			}
		}
	}
}

func (m *Manager) hasOpt(id int) bool {
	if id >= 0 && id < len(m.playerList) {
		player := m.GetPlayer(id)
		return player.canGang || player.canHu || player.canPen
	}else{
		for _, v := range m.playerList {
			if  v.canGang || v.canHu || v.canPen {
				return true
			}
		}
		return false
	}
}

func (m *Manager) initOptions(id int,card int,state int) {
	if state & Act_State_HU > 0 {
		m.checkHu(id,card)
	}

	if state & Act_State_AnG > 0 {
		m.checkAnGang(id)
	}

	if state & Act_State_WanG > 0 {
		m.checkWanGang(id)
	}

	if state & Act_State_DianG > 0 {
		m.checkDianGang(id,card)
	}

	if state & Act_State_Peng > 0 {
		m.checkPeng(id,card)
	}
}

func (m *Manager) dealOptions(card int){
	for k,v := range m.playerList {
		if v.canGang || v.canHu || v.canPen {
			v.canPen = false
			v.canGang = false
			v.gangCards = []int{}

			m.r.PushMessage(k,"game_action_push",protocol.ActionNotify{
				Id: k,
				Card: card,
				CanPeng: v.canPen,
				CanGang: v.canGang,
				CanDraw: v.canHu,
				CardGang: v.gangCards,
			})
		}
	}
}

func (m *Manager) notifyOptions(id int,card int) bool {
	flag := false
	for k,v := range m.playerList {
		if v.canGang || v.canHu || v.canPen {
			flag = true
			m.r.PushMessage(k,"game_action_push",protocol.ActionNotify{
				Id: k,
				Card: card,
				CanPeng: v.canPen,
				CanGang: v.canGang,
				CanDraw: v.canHu,
				CardGang: v.gangCards,
			})
		}
	}

	return flag
}

func (m *Manager)Shuffle(){
	rand.Seed(time.Now().UnixNano())
	for i:=m.cardIndex;i<len(m.cardList);i++{
		j := Rand(m.cardIndex,len(m.cardList)-1)
		m.cardList[i],m.cardList[j] = m.cardList[j],m.cardList[i]
	}

	// 配牌逻辑
	if CnfDebug {
		list :=	IntList(m.cardList)
		for _,v := range CardCnf {
			list.RemoveByValue(v)
		}
		for k,v := range CardCnf {
			list.Insert(v,k+1)
		}
	}
}

func (m *Manager)GetCardCount() int {
	return len(m.cardList) - m.cardIndex
}

func (m *Manager)Deal() int {
	if m.cardIndex < len(m.cardList) {
		card := m.cardList[m.cardIndex]
		m.cardIndex++
		//return card
		player := m.playerList[m.turn]
		player.holds = append(player.holds,card)
		player.cardMap[card]++
		return card
	}
	return -1
}

func (m *Manager)StartDeal(){
	var player *Player
	var card int
	for i:=0;i<4;i++{
		for j:=0;j<13;j++{
			player = m.playerList[i]
			card = m.cardList[m.cardIndex]
			player.holds = append(player.holds,card)
			player.cardMap[card]++
			m.cardIndex ++
		}
	}
	player = m.playerList[m.banker]
	card = m.cardList[m.cardIndex]
	player.holds = append(player.holds,card)
	player.cardMap[card]++
	m.cardIndex ++
}

func (m *Manager) StartGameMessage() {
	for i := 0; i< len(m.playerList); i++ {
		m.r.PushMessage(i,"game_holds_push",m.playerList[i].holds)
		m.r.PushMessage(i,"mj_count_push",m.GetCardCount())
		m.r.PushMessage(i,"game_num_push",1)
		m.r.PushMessage(i,"game_begin_push",m.banker)
		if m.config.HuanSanZhang {
			m.state = "huanpai"
			m.r.PushMessage(i,"game_huanpai_push",nil)
		}else{
			m.state = "dingque"
			m.r.PushMessage(i,"game_dingque_push",nil)
		}
	}
}

func (m *Manager)StartGame(){
	// 初始化数据
	m.init()
	m.Shuffle()
	m.StartDeal()
	m.StartGameMessage()
}

func (m *Manager) OnReady(id int){
	m.r.SetReady(id,true)
	player := m.GetPlayer(id)
	m.r.BroadcastEx("user_ready_push",protocol.ReadyResponse{
		UserId: player.Session.GetUserId(),
		Ready: true,
	},id);

	if !m.r.CheckReady() {
		return
	}
	m.StartGame()
}

func (m *Manager) MoveNext(id int) *Manager {
	if id >= 0 && id < 4 {
		m.turn = id
		return m
	}

	for {
		m.turn = (m.turn+1)%4
		if !m.GetPlayer(m.turn).drawed {
			return m
		}
	}
}

func (m *Manager) GainCard(){
	m.playCard = -1
	card := m.Deal()
	if card == -1 {
		//todo:游戏结束
		m.DoGameOver()
		return
	}

	m.r.Broadcast("mj_count_push",m.GetCardCount())
	m.r.PushMessage(m.turn,"game_mopai_push",card)
	player := m.GetPlayer(m.turn)
	player.canPlay = true
	m.r.Broadcast("game_chupai_push",player.Session.GetUserId())

	m.initOptions(m.turn,-1,Act_State_HU | Act_State_WanG | Act_State_AnG)
	m.notifyOptions(m.turn,card)


}

func (m *Manager) checkChanged() bool {
	for i:=0;i<len(m.playerList);i++ {
		if m.playerList[i].changeCards == nil {
			return false
		}
	}
	return true
}

func (m *Manager)OnChangeThree(id int,c1,c2,c3 int){
	if m.state != "huanpai" {
		log.Info("state err huanpai")
		return
	}

	player := m.playerList[id]
	if player.changeCards != nil {
		log.Info("has done change three")
		return
	}

	cards := []int{c1,c2,c3}
	if !player.checkCards(cards) {
		return
	}

	for _,c := range cards {
		player.holds = SliceRemoveByValue(player.holds,c)
		player.cardMap[c] = player.cardMap[c] - 1
	}
	player.changeCards = cards

	m.r.PushMessage(id,"game_holds_push",protocol.ChangeThreeResponse{
		UserId: player.Session.GetUserId(),
		ChangeCards: cards,
	})
	m.r.BroadcastEx("game_holds_push",protocol.ChangeThreeResponse{
		UserId: player.Session.GetUserId(),
		ChangeCards: []int{},
	},id)

	if m.checkChanged() {
		m.ChangeThreeCard()
	}
}

func (m *Manager)ChangeThreeCard(){
	var changeThreeFunc = func(player *Player,cards []int){
		for _,card := range cards {
			player.holds = append(player.holds, card)
			player.cardMap[card]++
		}
	}

	for i:=0;i<4;i++ {
		changeThreeFunc(m.playerList[i], m.playerList[(i+1)%4].changeCards)
	}
	m.state = "dingque"

	changeMethod := protocol.ChangeMethodNotify{Method: 1}
	for i:=0;i<len(m.playerList);i++ {
		m.r.PushMessage(i,"game_huanpai_over_push",changeMethod)
		m.r.PushMessage(i,"game_holds_push",m.playerList[i].holds)
	}
	m.r.Broadcast("game_dingque_push",nil)
}

func (m *Manager)checkLack() bool {
	for _,player := range m.playerList {
		if player.lack < 0 {
			return false
		}
	}
	return true
}

func (m *Manager) OnMakeLack(id int,lack int) {
	if lack < 0 || lack > 3 {
		return
	}

	if m.state != "dingque" {
		return
	}

	if m.playerList[id].lack > 0 {
		return
	}
	m.playerList[id].lack = lack

	if !m.checkLack() {
		m.r.Broadcast("game_dingque_notify_push",m.GetPlayer(id).Session.GetUserId())
		return
	}

	array := [4]int{}
	for i,v := range m.playerList {
		array[i] = v.lack
	}
	m.r.Broadcast("game_dingque_finish_push",array)
	m.r.Broadcast("game_playing_push",nil)
	//todo:检查所有玩家是否有可听牌的

	m.state = "playing"
	player := m.playerList[m.turn]
	player.canPlay = true
	m.r.Broadcast("game_chupai_push",player.Session.GetUserId())
	//todo:检查是否可杠、可胡
}

func (m *Manager) OnPlayCard (id int,card int) {
	log.Info("OnPlayCard card:%d",card)
	player := m.GetPlayer(id)
	if m.turn != id || player.drawed || !player.canPlay {
		return
	}

	if !player.checkCards(card) {
		return
	}

	player.holds = SliceRemoveByValue(player.holds,card)
	player.canPlay = false
	player.cardMap[card]--

	m.playCard = card
	m.playCnt++

	m.r.Broadcast("game_chupai_notify_push",protocol.PlayCardNotify{
		UserId: m.GetPlayerUserId(id),
		Card: card,
	})

	m.initOptions(id,card,Act_State_Peng | Act_State_DianG | Act_State_HU)
	if !m.notifyOptions(id,card) {
		m.r.Broadcast("guo_notify_push",protocol.PassNotify{
			UserId: m.GetPlayerUserId(id),
			Card: m.playCard,
		})
		m.MoveNext(-1).GainCard()
	}
}

func (m *Manager) othersCanHu(id int) bool {
	index := m.turn
	for {
		index = (index+1)%4
		if index == m.turn {
			return false
		}else{
			player := m.GetPlayer(index)
			if player.canHu && index != id {
				return true
			}
		}
	}
}

func (m *Manager) OnPeng(id int) {
	player := m.GetPlayer(id)
	if !player.canPen || player.drawed {
		return
	}

	if m.othersCanHu(id) {
		return
	}

	//if player.cardMap[m.playCard] < 2 {
	//	return
	//}
	player.holds = SliceRemoveByValue(player.holds,m.playCard)
	player.holds = SliceRemoveByValue(player.holds,m.playCard)
	player.cardMap[m.playCard] -= 2
	player.peCards = append(player.peCards,m.playCard)

	m.r.Broadcast("peng_notify_push",protocol.PassNotify{
		UserId: m.GetPlayerUserId(id),
		Card: m.playCard,
	})
	m.playCard = -1
	m.clearOpt(id)

	m.MoveNext(id)
	player.canPlay = true

	m.r.Broadcast("game_chupai_push",m.GetPlayerUserId(id))

}

func (m *Manager) DoGang(id int,gtype string,card int){
	player := m.GetPlayer(id)
	switch gtype {
	case "angang":
		player.agCards = append(player.agCards,card)
		player.holds = SliceRemoveByValue(player.holds,card,4)
		player.cardMap[card] -= 4
	case "diangang":
		player.dgCards = append(player.dgCards,card)
		player.holds = SliceRemoveByValue(player.holds,card,3)
		player.cardMap[card] -= 3
	case "wangang":
		player.wgCards = append(player.wgCards,card)
		player.peCards = SliceRemoveByValue(player.peCards,card)
		player.holds = SliceRemoveByValue(player.holds,card,1)
		player.cardMap[card] -= 1
	}

	m.r.Broadcast("gang_notify_push",protocol.GangNotify{
		UserId: m.GetPlayerUserId(id),
		Card: card,
		Type: gtype,
	})

	m.MoveNext(id).GainCard()
}

func (m *Manager) OnGang (id int,card int) {
	player := m.GetPlayer(id)
	if !player.canGang { //|| player.drawed {
		return
	}

	if m.othersCanHu(id) {
		return
	}

	var gtype string
	switch player.cardMap[card] {
	case 1:
		gtype = "wangang"
	case 3:
		gtype = "diangang"
	case 4:
		gtype = "angang"
	default:
		return
	}
	fmt.Println(gtype)
	m.clearOpt(-1)

	m.r.Broadcast("hangang_notify_push",id)

	//todo:是否可以抢杠

	m.DoGang(id,gtype,card)
}

func (m *Manager) OnPass(id int){
	player := m.GetPlayer(id)
	if !(player.canGang || player.canPen || player.canHu) {
		return
	}

	m.r.PushMessage(id,"guo_result",nil)
	m.clearOpt(id)

	if m.hasOpt(-1) {
		return
	}

	m.MoveNext(-1).GainCard()

}

func (m *Manager)OnHu(id int){
	player := m.GetPlayer(id)
	if player.drawed {
		return
	}
	player.drawed = true
	player.huList = append(player.huList,player.huCard)
	// todo:抢杠

	m.r.Broadcast("hu_push",protocol.HuNotify{
		Id: id,
		ZiMo: player.isZiMo,
		Card: player.huCard,
	})
	m.clearOpt(id)
	m.dealOptions(m.playCard)

	if m.othersCanHu(id) {
		return
	}

	if m.checkGameOver() {
		m.DoGameOver()
		return
	}

	m.clearOpt(-1)
	m.turn = id
	m.MoveNext(-1).GainCard()
	
}

func (m *Manager)checkGameOver() bool {
	cnt := 0
	for _,p := range m.playerList {
		if p.drawed {
			cnt++
		}
	}

	return cnt >= 3
}

func (m *Manager)DoGameOver(){
	players := make([]protocol.PlayerSettle,4)
	settles := make([]protocol.TotalSettle,4)

	m.r.ClearReady()

	for i:=0;i<4;i++ {
		p := m.playerList[i]
		data := &players[i]
		data.UserId = m.GetPlayerUserId(i)
		data.Peng = p.peCards
		data.Actions = []int{}
		data.WangGang = p.wgCards
		data.DianGang = p.dgCards
		data.AnGang = p.agCards
		data.NumOfGen = 0
		data.Holds = p.holds
		data.Fan = 0
		data.Score = 0
		data.TotalScore = 0
		data.HuInfo = []int{}

		data.Huorder = i

		settle := settles[i]
		settle.NumZiMo = p.numZiMo
		settle.NumJiaoPai = p.numJiePao
		settle.NumDianPao = p.numDianPao
		settle.NumAnGang = p.numAnGang
		settle.NumMingGang = p.numMingGang
		settle.NumChaDaJiao = p.numChaJiao
	}

	m.r.Broadcast("game_over_push",protocol.GameOverNotify{
		PlayerSettleList: players,
		TotalSettleList: settles,
	})
}