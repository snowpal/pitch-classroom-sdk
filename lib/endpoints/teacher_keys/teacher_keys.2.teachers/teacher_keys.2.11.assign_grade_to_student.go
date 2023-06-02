package teacherKeys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func AssignCourseGradeToStudent(
	jwtToken string,
	reqBody request.UpdateScaleValueReqBody,
	classroomParam request.ClassroomIdParam,
) (response.UpdateCourseScaleValue, error) {
	resCourseScaleValue := response.UpdateCourseScaleValue{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resCourseScaleValue, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		lib.RouteTeacherKeysAssignGradeToStudent,
		classroomParam.ResourceIds.CourseId,
		classroomParam.StudentId,
		classroomParam.ResourceIds.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resCourseScaleValue, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resCourseScaleValue, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resCourseScaleValue, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resCourseScaleValue, err
	}

	err = json.Unmarshal(body, &resCourseScaleValue)
	if err != nil {
		fmt.Println(err)
		return resCourseScaleValue, err
	}
	return resCourseScaleValue, nil
}
