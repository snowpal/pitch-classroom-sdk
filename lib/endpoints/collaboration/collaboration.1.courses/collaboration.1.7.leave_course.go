package collaboration

import (
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

func LeaveCourse(jwtToken string, courseParam common.ResourceIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteCollaborationLeaveCourse,
		courseParam.CourseId,
		courseParam.KeyId,
	)
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
