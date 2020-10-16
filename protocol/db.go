package protocol
type (
	FindUserRequest struct {
		UserId 		int64
		Account 	string
	}

	FindUserResponse struct {
		BaseResp
		UserModel
	}

	FindMessageRequset struct {
		Type 		string
		Version  	string
	}

	FindMessageResponse struct {
		BaseResp
		MessageModel
	}

	FindRoomsRequest struct {
		ServerId 	int
	}

	RoomModel struct {
		Uuid 		string
		Id 			string
		CreateTime 	int64
		NextButton	int64
		Config 		string
	}

	FindRoomsResponse struct {
		BaseResp
		RoomList 	[]RoomModel
	}
)