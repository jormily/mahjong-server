package logic

// 牌值
var (
	Card_Value_Start = 0
	Card_Value_End = 26
)
// 牌类型
var (
	Card_Type_Tong 	= 0
	Card_Type_Tiao 	= 1
	Card_Type_Wan	= 2
)
// 玩家操作
var (
	ACTION_CHUPAI = 1
	ACTION_MOPAI = 2
	ACTION_PENG = 3
	ACTION_GANG = 4
	ACTION_HU = 5
	ACTION_ZIMO = 6
)

var (
	Act_State_HU 	= 1
	Act_State_Peng 	= 2
	Act_State_Gang 	= 4
	Act_State_AnG	= 8
	Act_State_DianG = 16
	Act_State_WanG 	= 32
	Act_State_Play	= 64
)

