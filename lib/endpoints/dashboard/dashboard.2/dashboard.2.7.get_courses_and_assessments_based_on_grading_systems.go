package dashboard

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetCoursesAndAssessmentsBasedOnGradingSystems(jwtToken string) ([]response.GradingSystemsKey, error) {
	resGradingSystemsKeys := response.GradingSystemsKeys{}
	route, err := helpers2.GetRoute(lib.RouteDashboardGetCoursesAndAssessmentsBasedOnGradingSystems)
	if err != nil {
		return *resGradingSystemsKeys.Keys, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return *resGradingSystemsKeys.Keys, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return *resGradingSystemsKeys.Keys, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return *resGradingSystemsKeys.Keys, err
	}

	err = json.Unmarshal(body, &resGradingSystemsKeys)
	if err != nil {
		return *resGradingSystemsKeys.Keys, err
	}
	return *resGradingSystemsKeys.Keys, nil
}
