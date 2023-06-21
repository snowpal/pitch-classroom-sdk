package templates

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetAssessmentTemplates(jwtToken string) ([]response.AssessmentTemplate, error) {
	resAssessmentTemplates := response.AssessmentTemplates{}
	route, err := helpers2.GetRoute(lib.RouteTemplatesGetAssessmentTemplates)
	if err != nil {
		return resAssessmentTemplates.Templates, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resAssessmentTemplates.Templates, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resAssessmentTemplates.Templates, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resAssessmentTemplates.Templates, err
	}

	err = json.Unmarshal(body, &resAssessmentTemplates)
	if err != nil {
		return resAssessmentTemplates.Templates, err
	}
	return resAssessmentTemplates.Templates, nil
}
