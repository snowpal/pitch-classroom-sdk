package assessments

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
)

func MoveAssessment(jwtToken string, assessmentParam request.CopyMoveAssessmentParam) error {
	route, err := helpers.GetRoute(
		lib.RouteAssessmentsMoveAssessment,
		assessmentParam.AssessmentId,
		assessmentParam.KeyId,
		assessmentParam.CourseId,
		assessmentParam.TargetKeyId,
		assessmentParam.TargetCourseId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
