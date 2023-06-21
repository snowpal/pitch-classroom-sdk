package assessments

import (
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

func UnlinkAssessmentFromCourse(jwtToken string, assessmentParam common.ResourceIdParam) error {
	route, err := helpers.GetRoute(lib.RouteAssessmentsUnlinkAssessmentFromCourse, assessmentParam.CourseId, assessmentParam.AssessmentId, assessmentParam.KeyId)
	if err != nil {
		return err
	}

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
