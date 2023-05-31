package assessments

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

func UpdateAssessmentScaleValue(
	jwtToken string,
	reqBody request.UpdateScaleValueReqBody,
	assessmentParam common.ResourceIdParam,
) (response.UpdateAssessmentScaleValue, error) {
	resAssessmentscaleValue := response.UpdateAssessmentScaleValue{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resAssessmentscaleValue, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteAssessmentsUpdateAssessmentScaleValue,
		assessmentParam.AssessmentId,
		assessmentParam.KeyId,
		assessmentParam.CourseId,
	)
	if err != nil {
		fmt.Println(err)
		return resAssessmentscaleValue, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resAssessmentscaleValue, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAssessmentscaleValue, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAssessmentscaleValue, err
	}

	err = json.Unmarshal(body, &resAssessmentscaleValue)
	if err != nil {
		fmt.Println(err)
		return resAssessmentscaleValue, err
	}
	return resAssessmentscaleValue, nil
}
