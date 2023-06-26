package relations

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetRelationsForKey(jwtToken string, keyId string) (response.Relationships, error) {
	resRelations := response.Relations{}
	route, err := helpers2.GetRoute(lib.RouteRelationsGetRelationsForKey, keyId)
	if err != nil {
		return resRelations.Relationships, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resRelations.Relationships, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resRelations.Relationships, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resRelations.Relationships, err
	}

	err = json.Unmarshal(body, &resRelations)
	if err != nil {
		return resRelations.Relationships, err
	}
	return resRelations.Relationships, nil
}
