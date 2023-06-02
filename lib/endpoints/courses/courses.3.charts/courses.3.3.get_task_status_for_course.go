package courses

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetTaskStatusForCourse(jwtToken string, taskParam common.ResourceIdParam) (response.TasksStatusCourse, error) {
	resCourseTasksStatus := response.TasksStatusCourse{}
	route, err := helpers2.GetRoute(
		lib.RouteCoursesGetTaskStatusForCourse,
		taskParam.KeyId,
		taskParam.CourseId,
	)
	if err != nil {
		fmt.Println(err)
		return resCourseTasksStatus, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resCourseTasksStatus, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resCourseTasksStatus, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resCourseTasksStatus, err
	}

	err = json.Unmarshal(body, &resCourseTasksStatus)
	if err != nil {
		fmt.Println(err)
		return resCourseTasksStatus, err
	}
	return resCourseTasksStatus, nil
}
