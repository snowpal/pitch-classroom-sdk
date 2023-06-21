package assessments

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func UpdateAssessmentGrade(
	jwtToken string,
	reqBody request.UpdateGradeReqBody,
	assessmentParam common.ResourceIdParam,
) (response.UpdateAssessmentGrade, error) {
	resAssessmentGrade := response.UpdateAssessmentGrade{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resAssessmentGrade, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteAssessmentsUpdateAssessmentGrade,
		assessmentParam.AssessmentId,
		assessmentParam.KeyId,
		assessmentParam.CourseId,
	)
	if err != nil {
		return resAssessmentGrade, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		return resAssessmentGrade, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resAssessmentGrade, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resAssessmentGrade, err
	}

	err = json.Unmarshal(body, &resAssessmentGrade)
	if err != nil {
		return resAssessmentGrade, err
	}
	return resAssessmentGrade, nil
}
