package protocol

//type NewUserBroadcast struct {
//	Content string `json:"content"`
//}

type (
	EmptyRequest struct {

	}

	EmptyResponse struct {

	}

	BaseResp struct {
		Errcode 	int		`json:"errcode"`
		Errmsg 		string	`json:"errmsg"`
	}

	UserModel struct {
		Account 	string 	`json:"account"`
		UserId 		int64 	`json:"userid"`
		Name 		string 	`json:"name"`
		Level  		int64 	`json:"lv"`
		Exp   		int64 	`json:"exp"`
		Coins  		int64 	`json:"coins"`
		Gems 		int64 	`json:"gems"`
		Ip 			string 	`json:"ip"`
		Sex 		int64 	`json:"sex"`
		RoomId		string 	`json:"roomid"`
		Headimg 	string   `json:"headimgurl"`
	}

	MessageModel struct {
		Type    string `json:"type"`
		Msg     string `json:"msg"`
		Version string `json:"version"`
	}

	RoomCnf struct {
		Type 			string	`json:"type"`
		BaseScore		int		`json:"difen"`
		ZiMo			int		`json:"zimo"`
		JiangDui		bool	`json:"jiangdui"`
		HuanSanZhang 	bool	`json:"huansanzhang"`
		DianGangHua		string	`json:"dianganghua"`
		DianGanH		int64	`json:"diangangh"`
		MengQing		bool 	`json:"menqing"`
		TianDiHu		bool	`json:"tiandihu"`
		MaxFan			int 	`json:"zuidafanshu"`
		MaxGames		int		`json:"jushuxuanze"`
		Creator 		int		`json:"creator"`
	}

	Token struct {
		UserId 		int64	`json:"userId"`
		Time 		int64	`json:"time"`
		LifeTime	int64	`json:"lifeTime"`
	}
)