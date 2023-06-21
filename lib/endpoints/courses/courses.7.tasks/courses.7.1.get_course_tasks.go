package courses

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetCourseTasks(jwtToken string, taskParam request.TaskIdParam) ([]response.Task, error) {
	resTasks := response.Tasks{}
	route, err := helpers2.GetRoute(
		lib.RouteCoursesGetCourseTasks,
		*taskParam.CourseId,
		taskParam.KeyId,
	)
	if err != nil {
		return resTasks.Tasks, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resTasks.Tasks, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resTasks.Tasks, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resTasks.Tasks, err
	}

	err = json.Unmarshal(body, &resTasks)
	if err != nil {
		return resTasks.Tasks, err
	}
	return resTasks.Tasks, nil
}
