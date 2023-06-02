package courses

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

type CourseByTemplateParam struct {
	KeyId              string
	TemplateId         string
	ExcludeAssessments bool
	ExcludeTasks       bool
}

func AddCourseBasedOnTemplate(
	jwtToken string,
	reqBody request.AddCourseReqBody,
	courseParam CourseByTemplateParam,
) (response.Course, error) {
	resCourse := response.Course{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resCourse, err
	}
	payload := strings.NewReader(requestBody)
	var route string
	route, err = helpers2.GetRoute(
		lib.RouteCoursesAddCourseBasedOnTemplate,
		courseParam.KeyId,
		courseParam.TemplateId,
		strconv.FormatBool(courseParam.ExcludeAssessments),
		strconv.FormatBool(courseParam.ExcludeTasks),
	)
	if err != nil {
		fmt.Println(err)
		return resCourse, err
	}
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resCourse, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resCourse, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resCourse, err
	}

	err = json.Unmarshal(body, &resCourse)
	if err != nil {
		fmt.Println(err)
		return resCourse, err
	}
	return resCourse, nil
}
