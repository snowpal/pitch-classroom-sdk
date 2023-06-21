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

func GetCourseCollaborators(jwtToken string, courseParam common.ResourceIdParam) (response.Course, error) {
	resCourse := response.Course{}
	route, err := helpers2.GetRoute(
		lib.RouteCollaborationGetCourseCollaborators,
		courseParam.CourseId,
		courseParam.KeyId,
	)
	if err != nil {
		return resCourse, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resCourse, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resCourse, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resCourse, err
	}

	err = json.Unmarshal(body, &resCourse)
	if err != nil {
		return resCourse, err
	}
	return resCourse, nil
}
