package relations

import (
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
)

func UnrelateCourseFromCourse(jwtToken string, relationParam request.CourseToCourseRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsUnrelateCourseFromCourse,
		relationParam.CourseId,
		relationParam.TargetCourseId,
	)
	if err != nil {
		return err
	}

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
