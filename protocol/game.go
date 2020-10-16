package protocol

type EnterTokenResponse struct {
	BaseResp
	Tok 	Token	`json:"token"`
}

type EmptyMsg struct {

}

type LoginMsg struct {
	Token 	string 	`json:"token"`
	RoomId	string 	`json:"roomid"`
	Time	int64	`json:"time"`
	Sign 	string 	`json:"sign"`
}

type PlayerInfo struct {
	UserId 	int64	`json:"userid"`
	Ip 		string	`json:"ip"`
	Name 	string	`json:"name"`
	Score 	int		`json:"score"`
	Online	bool	`json:"online"`
	Ready	bool	`json:"ready"`
	SeateIndex	int `json:"seatindex"`
}

type LoginResult struct {
	RoomId 		string 			`json:"roomid"`
	Conf 		RoomCnf			`json:"conf"`
	GameCount	int				`json:"numofgames"`
	Seats		[]PlayerInfo 	`json:"seats"`
}

type LoginResultResponse struct {
	BaseResp
	Result 	LoginResult 	`json:"data"`
}

type ReadyResponse struct {
	UserId 	int64	`json:"userid"`
	Ready 	bool 	`json:"ready"`
}

type ChangeThreeResponse struct {
	UserId 	int64	`json:"userid"`
	ChangeCards	[]int	`json:"huanpais"`
}

type ChangeMethodNotify struct {
	Method 	int		`json:"method"`
}

type MakeLackNotify struct {

}

type PlayCardNotify struct {
	Card 	int		`json:"pai"`
	UserId 	int64	`json:"userId"`
}

type ActionNotify struct {
	Id 			int 	`json:"si"`
	Card 		int 	`json:"pai"`
	CanDraw		bool	`json:"hu"`
	CanPeng		bool	`json:"peng"`
	CanGang		bool	`json:"gang"`
	CardGang	[]int 	`json:"gangpai"`
}

type PassNotify struct {
	UserId 	int64	`json:"userId"`
	Card 	int 	`json:"pai"`
}

type GangNotify struct {
	UserId 	int64	`json:"userId"`
	Card 	int 	`json:"pai"`
	Type 	string 	`json:"gangtype"`
}

type HuNotify struct {
	Id 		int		`json:"seatindex"`
	ZiMo	bool	`json:"iszimo"`
	Card 	int 	`json:"hupai"`
}

type PlayerSettle struct {
	UserId 		int64 	`json:"userId"`
	Peng		[]int	`json:"pengs"`
	Actions 	[]int 	`json:"actions"`
	WangGang	[]int	`json:"wangangs"`
	DianGang	[]int	`json:"diangangs"`
	AnGang		[]int 	`json:"angangs"`
	NumOfGen	int 	`json:"numofgen"`
	Holds		[]int	`json:"holds"`
	Fan 		int 	`json:"fan"`
	Score 		int 	`json:"score"`
	TotalScore	int		`json:"totalscore"`
	QingYiSe	bool 	`json:"qingyise"`
	Pattern 	string 	`json:"pattern"`
	IsGangHu	bool	`json:"isganghu"`
	MenQing		bool	`json:"menqing"`
	ZhongZhang 	bool	`json:"zhongzhang"`

	JingGouHu	bool	`json:"jingouhu"`
	HaiDiHu		bool	`json:"haidihu"`
	TianHu		bool	`json:"tianhu"`
	DiHu		bool	`json:"dihu"`
	Huorder		int		`json:"huorder"`
	HuInfo 		[]int 	`json:"huinfo"`  //todo:完善
}

type TotalSettle struct {
	NumZiMo 	int 	`json:"numzimo"`
	NumJiaoPai	int		`json:"numjiaopai"`
	NumDianPao	int 	`json:"numdianpao"`
	NumAnGang 	int 	`json:"numangang"`
	NumMingGang	int		`json:"numminggang"`
	NumChaDaJiao int 	`json:"numchadajiao"`
}

type GameOverNotify struct {
	PlayerSettleList 	[]PlayerSettle 		`json:"results"`
	TotalSettleList		[]TotalSettle		`json:"endinfo"`
	IsOver 				bool 				`json:"isOver"`
}
