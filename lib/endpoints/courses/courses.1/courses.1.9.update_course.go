package courses

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

type UpdateCourseReqBody struct {
	Name              *string `json:"courseName"`
	CourseId          *string `json:"courseId"`
	SimpleDescription *string `json:"simpleDescription"`
	DueDate           *string `json:"courseDueDate"`
	StartTime         *string `json:"courseStartTime"`
	EndTime           *string `json:"courseEndTime"`
	Color             *string `json:"courseColor"`
	Tags              *string `json:"courseTags"`
	KanbanMode        *bool   `json:"kanbanMode"`
	Completed         bool    `json:"courseCompleted"`
}

func UpdateCourse(jwtToken string, reqBody UpdateCourseReqBody, courseParam common.ResourceIdParam) (response.Course, error) {
	resCourse := response.Course{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resCourse, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteCoursesUpdateCourse, courseParam.CourseId, courseParam.KeyId)
	if err != nil {
		return resCourse, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
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
