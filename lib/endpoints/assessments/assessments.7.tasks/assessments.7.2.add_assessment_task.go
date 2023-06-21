package assessments

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func AddAssessmentTask(
	jwtToken string,
	reqBody request.AddTaskReqBody,
	taskParam request.TaskIdParam,
) (response.Task, error) {
	resTask := response.Task{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resTask, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteAssessmentsAddAssessmentTask,
		*taskParam.AssessmentId,
		taskParam.KeyId,
		*taskParam.CourseId,
	)
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		return resTask, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resTask, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resTask, err
	}

	err = json.Unmarshal(body, &resTask)
	if err != nil {
		return resTask, err
	}
	return resTask, nil
}
