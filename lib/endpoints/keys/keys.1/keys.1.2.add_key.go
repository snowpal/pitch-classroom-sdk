package keys

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func AddKey(jwtToken string, reqBody request.AddKeyReqBody) (response.Key, error) {
	resKey := response.Key{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resKey, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteKeysAddKey)
	if err != nil {
		return resKey, err
	}
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		return resKey, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resKey, err
	}

	defer helpers2.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		return resKey, err
	}

	err = json.Unmarshal(body, &resKey)
	if err != nil {
		return resKey, err
	}
	return resKey, nil
}
