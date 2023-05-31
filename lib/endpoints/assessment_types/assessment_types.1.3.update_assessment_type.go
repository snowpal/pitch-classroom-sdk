package assessmentTypes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func UpdateAssessmentType(jwtToken string, reqBody request.AssessmentTypeReqBody, assessmentTypeId string) (response.AssessmentType, error) {
	resAssessmentType := response.AssessmentType{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resAssessmentType, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteAssessmentTypesUpdateAssessmentType, assessmentTypeId)
	if err != nil {
		fmt.Println(err)
		return resAssessmentType, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resAssessmentType, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAssessmentType, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAssessmentType, err
	}

	err = json.Unmarshal(body, &resAssessmentType)
	if err != nil {
		fmt.Println(err)
		return resAssessmentType, err
	}
	return resAssessmentType, nil
}
