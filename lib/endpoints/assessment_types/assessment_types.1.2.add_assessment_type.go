package assessmentTypes

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

func AddAssessmentType(jwtToken string, reqBody request.AssessmentTypeReqBody) (response.AssessmentType, error) {
	resAssessmentType := response.AssessmentType{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resAssessmentType, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteAssessmentTypesAddAssessmentType)
	if err != nil {
		return resAssessmentType, err
	}

	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		return resAssessmentType, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resAssessmentType, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resAssessmentType, err
	}

	err = json.Unmarshal(body, &resAssessmentType)
	if err != nil {
		return resAssessmentType, err
	}
	return resAssessmentType, nil
}
