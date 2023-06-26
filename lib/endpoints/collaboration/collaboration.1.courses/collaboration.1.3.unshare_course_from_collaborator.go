package collaboration

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func UnshareCourseFromCollaborator(jwtToken string, courseAclParam common.AclParam) (response.Course, error) {
	resCourse := response.Course{}
	route, err := helpers2.GetRoute(
		lib.RouteCollaborationUnshareCourseFromCollaborator,
		courseAclParam.ResourceIds.CourseId,
		courseAclParam.UserId,
		courseAclParam.ResourceIds.KeyId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		return resCourse, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resCourse, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resCourse, err
	}

	err = json.Unmarshal(body, &resCourse)
	if err != nil {
		return resCourse, err
	}
	return resCourse, nil
}
