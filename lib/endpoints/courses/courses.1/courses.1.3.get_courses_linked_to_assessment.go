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

func GetCoursesLinkedToAssessment(jwtToken string, courseParam common.ResourceIdParam) ([]response.Course, error) {
	resCourses := response.Courses{}
	route, err := helpers2.GetRoute(
		lib.RouteCoursesGetCoursesLinkedToAssessment,
		courseParam.AssessmentId,
		courseParam.KeyId,
	)
	if err != nil {
		return resCourses.Courses, err
	}

	req, _ := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resCourses.Courses, err
	}

	helpers2.AddUserHeaders(jwtToken, req)
	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resCourses.Courses, err
	}

	defer helpers2.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		return resCourses.Courses, err
	}

	err = json.Unmarshal(body, &resCourses)
	if err != nil {
		return resCourses.Courses, err
	}
	return resCourses.Courses, nil
}
