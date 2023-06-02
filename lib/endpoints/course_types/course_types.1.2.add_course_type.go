package course_types

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func AddCourseType(jwtToken string, reqBody request.CourseTypeReqBody) (response.CourseType, error) {
	resCourseType := response.CourseType{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resCourseType, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteCourseTypesAddCourseType)
	if err != nil {
		fmt.Println(err)
		return resCourseType, err
	}

	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resCourseType, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resCourseType, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resCourseType, err
	}

	err = json.Unmarshal(body, &resCourseType)
	if err != nil {
		fmt.Println(err)
		return resCourseType, err
	}
	return resCourseType, nil
}
