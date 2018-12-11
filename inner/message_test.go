package inner

import "testing"

func Test_Message_Send(t *testing.T) {

	req := &MessageSendReq{AgentId:1000005,ToUser: "john", MsgType: "text", Text: &MessageText{Content: "Always Hello World!!!"}}

	res, err := Message_Send(req)

	if err != nil {

		t.Fatal(err)
		return
	}

	t.Log("res",res)
}
