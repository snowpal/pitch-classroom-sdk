package assessments

import (
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
)

func DeleteAssessmentComment(jwtToken string, commentParam request.CommentIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteAssessmentsDeleteAssessmentComment,
		*commentParam.CommentId,
		commentParam.KeyId,
		*commentParam.CourseId,
		*commentParam.AssessmentId,
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
