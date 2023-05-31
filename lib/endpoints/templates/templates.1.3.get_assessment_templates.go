package templates

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetAssessmentTemplates(jwtToken string) ([]response.AssessmentTemplate, error) {
	resPodTemplates := response.AssessmentTemplates{}
	route, err := helpers2.GetRoute(lib.RouteTemplatesGetAssessmentTemplates)
	if err != nil {
		fmt.Println(err)
		return resPodTemplates.Templates, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resPodTemplates.Templates, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPodTemplates.Templates, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPodTemplates.Templates, err
	}

	err = json.Unmarshal(body, &resPodTemplates)
	if err != nil {
		fmt.Println(err)
		return resPodTemplates.Templates, err
	}
	return resPodTemplates.Templates, nil
}
