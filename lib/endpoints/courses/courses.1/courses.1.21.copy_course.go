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

func CopyCourse(jwtToken string, courseParam request.CopyMoveCourseParam) (response.Course, error) {
	var resCourse response.Course
	route, err := helpers2.GetRoute(
		lib.RouteCoursesCopyCourse,
		courseParam.CourseId,
		courseParam.KeyId,
		strconv.FormatBool(courseParam.AllTasks),
		strings.Join(courseParam.AssessmentIds, ","),
		strconv.FormatBool(courseParam.AllAssessments),
		strconv.FormatBool(courseParam.AllChecklists),
		courseParam.TargetKeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resCourse, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, nil)
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
