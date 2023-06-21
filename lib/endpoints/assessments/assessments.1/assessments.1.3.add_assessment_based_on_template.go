package assessments

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func AddAssessmentBasedOnTemplate(
	jwtToken string,
	reqBody request.AddAssessmentReqBody,
	assessmentParam request.AssessmentByTemplateParam,
) (response.Assessment, error) {
	resAssessment := response.Assessment{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resAssessment, err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers2.GetRoute(
		lib.RouteAssessmentsAddAssessmentBasedOnTemplate,
		*assessmentParam.CourseId,
		assessmentParam.KeyId,
		assessmentParam.TemplateId,
		strconv.FormatBool(assessmentParam.ExcludeTasks),
	)
	if err != nil {
		return resAssessment, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		return resAssessment, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resAssessment, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resAssessment, err
	}

	err = json.Unmarshal(body, &resAssessment)
	if err != nil {
		return resAssessment, err
	}
	return resAssessment, err
}
