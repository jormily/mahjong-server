package protocol

type (
	UserStatusRequest struct {
		Account string 		`json:"account"`
	}

	UserStatusResponse struct {
		BaseResp
		Gems 	int 		`json:"gems"`
	}

	MessageRequest struct {
		Type 	string 		`json:"type"`
		Version string		`json:"version"`
	}

	MessageResponse struct {
		BaseResp
		MessageModel
	}

	BaseInfoRequset struct {
		UserId 	int64 		`json:"userid"`
	}

	BaseInfoResponse struct {
		BaseResp
		Name 		string	`json:"name"`
		Sex 		int 	`json:"sex"`
		HeadImgUrl	string	`json:"headimgurl"`
	}
)
