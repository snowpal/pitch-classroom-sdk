package courses

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func UpdateCourseGrade(
	jwtToken string,
	reqBody request.UpdateGradeReqBody,
	courseParam common.ResourceIdParam,
) (response.UpdateCourseGrade, error) {
	resCourseGrade := response.UpdateCourseGrade{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resCourseGrade, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteCoursesUpdateCourseGrade,
		courseParam.CourseId,
		courseParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resCourseGrade, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resCourseGrade, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resCourseGrade, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resCourseGrade, err
	}

	err = json.Unmarshal(body, &resCourseGrade)
	if err != nil {
		fmt.Println(err)
		return resCourseGrade, err
	}
	return resCourseGrade, nil
}
