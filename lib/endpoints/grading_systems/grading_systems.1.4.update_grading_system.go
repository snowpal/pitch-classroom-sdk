package gradingSystems

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
)

func UpdateGradingSystem(jwtToken string, reqBody request.GradingSystemReqBody, gradingSystemId string) (response.GradingSystem, error) {
	resGradingSystem := response.GradingSystem{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resGradingSystem, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteGradingSystemsUpdateGradingSystem, gradingSystemId)
	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resGradingSystem, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resGradingSystem, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resGradingSystem, err
	}

	err = json.Unmarshal(body, &resGradingSystem)
	if err != nil {
		fmt.Println(err)
		return resGradingSystem, err
	}
	return resGradingSystem, nil
}
