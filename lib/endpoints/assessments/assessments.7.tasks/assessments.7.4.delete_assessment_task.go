package assessments

import (
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
)

func DeleteAssessmentTask(jwtToken string, taskParam request.TaskIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteAssessmentsDeleteAssessmentTask,
		*taskParam.TaskId,
		taskParam.KeyId,
		*taskParam.CourseId,
		*taskParam.AssessmentId,
	)
	req, err := http.NewRequest(http.MethodDelete, route, nil)
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
