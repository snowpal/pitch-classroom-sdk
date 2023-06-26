package keys

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
)

func GetGradingSystemsAlongWithCoursesAndAssessmentsBasedOnThem(jwtToken string, keyId string) (response.GradingSystemsKey, error) {
	resGradingSystemsKey := response.GradingSystemsKey{}
	route, err := helpers2.GetRoute(lib.RouteKeysGetGradingSystemsAlongWithCoursesAndAssessmentsBasedOnThem, keyId)
	if err != nil {
		return resGradingSystemsKey, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resGradingSystemsKey, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resGradingSystemsKey, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resGradingSystemsKey, err
	}

	err = json.Unmarshal(body, &resGradingSystemsKey)
	if err != nil {
		return resGradingSystemsKey, err
	}
	return resGradingSystemsKey, nil
}
