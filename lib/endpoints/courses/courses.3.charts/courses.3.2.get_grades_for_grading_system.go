package courses

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
)

func GetGradesForGradingSystem(jwtToken string, gradingSystemParam request.GradingSystemIdParam) (response.Grades, error) {
	resGrades := response.Grades{}
	route, err := helpers2.GetRoute(
		lib.RouteCoursesGetGradesForGradingSystem,
		gradingSystemParam.KeyId,
		*gradingSystemParam.CourseId,
		gradingSystemParam.GradingSystemId,
	)
	if err != nil {
		return resGrades, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resGrades, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resGrades, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resGrades, err
	}

	err = json.Unmarshal(body, &resGrades)
	if err != nil {
		return resGrades, err
	}
	return resGrades, nil
}
