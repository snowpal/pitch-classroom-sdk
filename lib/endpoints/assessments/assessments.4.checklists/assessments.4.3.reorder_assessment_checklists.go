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

func ReorderAssessmentChecklists(
	jwtToken string,
	reqBody request.ReorderChecklistsReqBody,
	checklistParam request.ChecklistIdParam,
) ([]response.Checklist, error) {
	resChecklists := response.Checklists{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resChecklists.Checklists, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteAssessmentsReorderAssessmentChecklists,
		*checklistParam.AssessmentId,
		checklistParam.KeyId,
		*checklistParam.CourseId,
	)
	if err != nil {
		return resChecklists.Checklists, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		return resChecklists.Checklists, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resChecklists.Checklists, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resChecklists.Checklists, err
	}

	err = json.Unmarshal(body, &resChecklists)
	if err != nil {
		return resChecklists.Checklists, err
	}
	return resChecklists.Checklists, nil
}
