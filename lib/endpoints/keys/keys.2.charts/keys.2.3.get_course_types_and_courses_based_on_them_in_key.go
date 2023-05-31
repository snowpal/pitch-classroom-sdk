package keys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetCourseTypesAndCoursesBasedOnThemInKey(jwtToken string, keyId string) (response.CourseTypesKey, error) {
	resBlockTypesKey := response.CourseTypesKey{}
	route, err := helpers2.GetRoute(lib.RouteKeysGetCourseTypesAndCoursesBasedOnThemInKey, keyId)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKey, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKey, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKey, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKey, err
	}

	err = json.Unmarshal(body, &resBlockTypesKey)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKey, err
	}
	return resBlockTypesKey, nil
}
