package keys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
)

func GetCourseGrades(jwtToken string, gradingSystemParam request.GradingSystemIdParam) (response.Grades, error) {
	resGrades := response.Grades{}
	route, err := helpers2.GetRoute(
		lib.RouteKeysGetCourseGrades,
		gradingSystemParam.KeyId,
		gradingSystemParam.GradingSystemId,
	)
	if err != nil {
		fmt.Println(err)
		return resGrades, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resGrades, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resGrades, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resGrades, err
	}

	err = json.Unmarshal(body, &resGrades)
	if err != nil {
		fmt.Println(err)
		return resGrades, err
	}
	return resGrades, nil
}
