package keys

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetCourseTypesAndCoursesBasedOnThemInKey(jwtToken string, keyId string) (response.CourseTypesKey, error) {
	resCourseTypesKey := response.CourseTypesKey{}
	route, err := helpers2.GetRoute(lib.RouteKeysGetCourseTypesAndCoursesBasedOnThemInKey, keyId)
	if err != nil {
		return resCourseTypesKey, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resCourseTypesKey, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resCourseTypesKey, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resCourseTypesKey, err
	}

	err = json.Unmarshal(body, &resCourseTypesKey)
	if err != nil {
		return resCourseTypesKey, err
	}
	return resCourseTypesKey, nil
}
