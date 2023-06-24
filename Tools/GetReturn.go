package Tools

import (
	"wolWebUi/Model"
)

func GetFailMsg(msg string) Model.RetMsg {
	var retmsg Model.RetMsg
	retmsg.Count = 0
	retmsg.Code = -1
	retmsg.Data = nil
	retmsg.Msg = msg

	return retmsg
}

func GetSuccMsg(count int, data interface{}) Model.RetMsg {
	var retmsg Model.RetMsg
	retmsg.Count = count
	retmsg.Code = 0
	retmsg.Data = data
	retmsg.Msg = ""

	return retmsg
}
