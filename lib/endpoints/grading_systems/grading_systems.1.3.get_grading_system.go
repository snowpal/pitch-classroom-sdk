package grading_systems

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
)

func GetGradingSystem(jwtToken string, gradingSystemId string) (response.GradingSystem, error) {
	resGradingSystem := response.GradingSystem{}
	route, err := helpers2.GetRoute(lib.RouteGradingSystemsGetGradingSystem, gradingSystemId)
	req, err := http.NewRequest(http.MethodGet, route, nil)
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
