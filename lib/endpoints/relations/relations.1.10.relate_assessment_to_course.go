package relations

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
)

func relateCourseToAssessment(jwtToken string, route string) error {
	req, err := http.NewRequest(http.MethodPatch, route, nil)
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

func RelateCourseToAssessment(jwtToken string, relationParam request.CourseToAssessmentRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsRelateAssessmentToCourse,
		relationParam.CourseId,
		relationParam.TargetAssessmentId,
		relationParam.TargetKeyId,
		relationParam.TargetCourseId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err != nil {
		fmt.Println(err)
		return err
	}
	err = relateCourseToAssessment(jwtToken, route)
	if err != nil {
		return err
	}
	return nil
}
