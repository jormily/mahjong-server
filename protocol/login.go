package protocol

type (
	ServerInfoResponse struct {
		Version int 	`json:"version"`
		Hall  	string	`json:"hall"`
		Appweb 	string	`json:"appweb"`
	}


	GuestAuthRequest struct {
		Account 	string 	`json:"account"`
	}

	GuestAuthResponse struct {
		BaseResp
		Account 	string 	`json:"account"`
		Sign 		string 	`json:"sign"`
	}

	LoginRequest struct {
		Account 	string 	`json:"account"`
	}

	LoginResponse struct {
		BaseResp
		UserModel
	}

	BaseInfoResquest struct {
		UserId 		int64 	`json:"userid"`
	}

)