package assessments

import (
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
)

func AddGradingSystemToAssessment(jwtToken string, assessmentParam request.GradingSystemIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteAssessmentsAddGradingSystemToAssessment,
		*assessmentParam.AssessmentId,
		assessmentParam.GradingSystemId,
		assessmentParam.KeyId,
		*assessmentParam.CourseId,
	)
	if err != nil {
		return err
	}
	println(route)

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		return err
	}
	return nil
}
