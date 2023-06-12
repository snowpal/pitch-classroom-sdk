package gradingSystems

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
)

func GetAssessmentsUsingGradingSystem(jwtToken string, gradingSystemId string) ([]response.Assessment, error) {
	resAssessments := response.Assessments{}
	route, err := helpers2.GetRoute(lib.RouteGradingSystemsGetAssessmentsUsingGradingSystem, gradingSystemId)
	if err != nil {
		fmt.Println(err)
		return resAssessments.Assessments, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAssessments.Assessments, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAssessments.Assessments, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAssessments.Assessments, err
	}

	err = json.Unmarshal(body, &resAssessments)
	if err != nil {
		fmt.Println(err)
		return resAssessments.Assessments, err
	}
	return resAssessments.Assessments, nil
}
