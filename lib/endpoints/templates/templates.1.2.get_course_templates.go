package templates

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetCourseTemplates(jwtToken string) ([]response.CourseTemplate, error) {
	resCourseTemplates := response.CourseTemplates{}
	route, err := helpers2.GetRoute(lib.RouteTemplatesGetCourseTemplates)
	if err != nil {
		return resCourseTemplates.Templates, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resCourseTemplates.Templates, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resCourseTemplates.Templates, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resCourseTemplates.Templates, err
	}

	err = json.Unmarshal(body, &resCourseTemplates)
	if err != nil {
		return resCourseTemplates.Templates, err
	}
	return resCourseTemplates.Templates, nil
}
