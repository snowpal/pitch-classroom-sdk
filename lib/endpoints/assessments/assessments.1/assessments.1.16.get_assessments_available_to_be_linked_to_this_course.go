package assessments

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetAssessmentsAvailableToBeLinked(jwtToken string, assessmentParam common.ResourceIdParam) ([]response.Assessment, error) {
	resAssessments := response.Assessments{}
	route, err := helpers2.GetRoute(lib.RouteAssessmentsGetAssessmentsAvailableToBeLinkedToThisCourse, assessmentParam.CourseId, assessmentParam.KeyId)
	if err != nil {
		return resAssessments.Assessments, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resAssessments.Assessments, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resAssessments.Assessments, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resAssessments.Assessments, err
	}

	err = json.Unmarshal(body, &resAssessments)
	if err != nil {
		return resAssessments.Assessments, err
	}
	return resAssessments.Assessments, nil
}
