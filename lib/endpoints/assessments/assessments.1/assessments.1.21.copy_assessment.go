package assessments

import (
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
)

func CopyAssessment(jwtToken string, assessmentParam request.CopyMoveAssessmentParam) error {
	route, err := helpers.GetRoute(
		lib.RouteAssessmentsCopyAssessment,
		assessmentParam.AssessmentId,
		assessmentParam.KeyId,
		assessmentParam.CourseId,
		strconv.FormatBool(assessmentParam.AllTasks),
		strconv.FormatBool(assessmentParam.AllChecklists),
		assessmentParam.TargetKeyId,
		assessmentParam.TargetCourseId,
	)
	if err != nil {
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, nil)
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
