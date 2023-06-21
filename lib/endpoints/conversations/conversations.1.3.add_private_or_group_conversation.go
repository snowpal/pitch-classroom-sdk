package conversations

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
)

type ConversationReqBody struct {
	MessageText string `json:"messageText"`
	Usernames   string `json:"userNames"`
}

func AddPrivateOrGroupConversation(jwtToken string, reqBody ConversationReqBody) (response.Conversation, error) {
	resConversation := response.Conversation{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resConversation, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteConversationsAddPrivateOrGroupConversation)
	if err != nil {
		return resConversation, err
	}

	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		return resConversation, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resConversation, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resConversation, err
	}

	err = json.Unmarshal(body, &resConversation)
	if err != nil {
		return resConversation, err
	}
	return resConversation, nil
}
