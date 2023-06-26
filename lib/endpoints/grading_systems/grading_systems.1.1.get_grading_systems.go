package grading_systems

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
)

type GetGradingSystemsParam struct {
	IncludeCounts bool
	ExcludeEmpty  bool
}

func GetGradingSystems(jwtToken string, gradingSystemsParam GetGradingSystemsParam) ([]response.GradingSystem, error) {
	resGradingSystems := response.GradingSystems{}
	route, err := helpers2.GetRoute(
		lib.RouteGradingSystemsGetGradingSystems,
		strconv.FormatBool(gradingSystemsParam.IncludeCounts),
		strconv.FormatBool(gradingSystemsParam.ExcludeEmpty),
	)
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resGradingSystems.GradingSystems, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resGradingSystems.GradingSystems, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resGradingSystems.GradingSystems, err
	}

	err = json.Unmarshal(body, &resGradingSystems)
	if err != nil {
		return resGradingSystems.GradingSystems, err
	}
	return resGradingSystems.GradingSystems, nil
}
