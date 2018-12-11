package inner

type MessageText struct {
	Content string `json:"content"`
}

type MessageSendReq struct {
	ToUser  string       `json:"touser" `
	ToParty string       `json:"toparty"`
	ToTag   string       `json:"totag"`
	MsgType string       `json:"msgtype"`
	AgentId int          `json:"agentid"`
	Text    *MessageText `json:"text"`
	Safe    int          `json:"safe"`
}

type MessageSendRes struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
}

func Message_Send(req *MessageSendReq) (res *MessageSendRes, err error) {

	err = postBody("/message/send", req, &res)

	return
}
