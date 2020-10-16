package logic

type Room interface {
	PushMessage(index int,route string,msg interface{})
	Broadcast(route string,msg interface{})
	BroadcastEx(route string,msg interface{},id int)
	SetReady(id int,ready bool)
	ClearReady()
	CheckReady() bool
}
