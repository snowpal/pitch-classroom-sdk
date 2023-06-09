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

func GetAssessment(jwtToken string, assessmentParam common.ResourceIdParam) (response.Assessment, error) {
	resAssessment := response.Assessment{}
	route, err := helpers2.GetRoute(lib.RouteAssessmentsGetAssessment, assessmentParam.AssessmentId, assessmentParam.KeyId, assessmentParam.CourseId)
	if err != nil {
		return resAssessment, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resAssessment, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resAssessment, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resAssessment, err
	}

	err = json.Unmarshal(body, &resAssessment)
	if err != nil {
		return resAssessment, err
	}
	return resAssessment, nil
}
