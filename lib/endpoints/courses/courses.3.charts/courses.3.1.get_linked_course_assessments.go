package courses

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetLinkedAssessments(jwtToken string, courseParam common.ResourceIdParam) (response.LinkedResources, error) {
	resLinkedAssessments := response.LinkedResources{}
	route, err := helpers2.GetRoute(
		lib.RouteCoursesGetLinkedAssessments,
		courseParam.KeyId,
		courseParam.CourseId,
	)
	if err != nil {
		return resLinkedAssessments, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resLinkedAssessments, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resLinkedAssessments, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resLinkedAssessments, err
	}

	err = json.Unmarshal(body, &resLinkedAssessments)
	if err != nil {
		return resLinkedAssessments, err
	}
	return resLinkedAssessments, nil
}
