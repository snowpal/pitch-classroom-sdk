package teacherKeys

import (
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
)

func BulkPublishAssessmentGradesForStudents(
	jwtToken string,
	reqBody request.PublishGradesReqBody,
	assessmentParam common.ResourceIdParam,
) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		return err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		lib.RouteTeacherKeysBulkPublishAssessmentGradesForStudents,
		assessmentParam.AssessmentId,
		assessmentParam.KeyId,
		assessmentParam.CourseId,
	)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
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
