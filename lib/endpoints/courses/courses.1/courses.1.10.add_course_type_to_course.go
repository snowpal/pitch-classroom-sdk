package courses

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
)

type AddBlockTypeIdParam struct {
	KeyId       string
	CourseId    string
	BlockTypeId string
}

func AddCourseTypeToCourse(jwtToken string, assessmentParam AddBlockTypeIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteCoursesAddCourseTypeToCourse,
		assessmentParam.CourseId,
		assessmentParam.BlockTypeId,
		assessmentParam.KeyId,
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
