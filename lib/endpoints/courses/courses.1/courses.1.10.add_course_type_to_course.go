package courses

import (
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
)

type AddCourseTypeIdParam struct {
	KeyId        string
	CourseId     string
	CourseTypeId string
}

func AddCourseTypeToCourse(jwtToken string, assessmentParam AddCourseTypeIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteCoursesAddCourseTypeToCourse,
		assessmentParam.CourseId,
		assessmentParam.CourseTypeId,
		assessmentParam.KeyId,
	)
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
