package relations

import (
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
)

func relateAssessmentToAssessment(jwtToken string, route string) error {
	req, err := http.NewRequest(http.MethodPatch, route, nil)
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

func RelateAssessmentToAssessment(jwtToken string, relationParam request.AssessmentToAssessmentRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsRelateAssessmentToAssessment,
		relationParam.AssessmentId,
		relationParam.SourceKeyId,
		relationParam.SourceCourseId,
		relationParam.TargetAssessmentId,
		relationParam.TargetKeyId,
		relationParam.TargetCourseId,
	)
	if err != nil {
		return err
	}
	err = relateAssessmentToAssessment(jwtToken, route)
	if err != nil {
		return err
	}
	return nil
}
