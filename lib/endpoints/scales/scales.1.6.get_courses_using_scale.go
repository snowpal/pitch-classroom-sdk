package scales

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetCoursesUsingScale(jwtToken string, scaleId string) ([]response.Course, error) {
	resCourses := response.Courses{}
	route, err := helpers2.GetRoute(lib.RouteScalesGetCoursesUsingScale, scaleId)
	if err != nil {
		fmt.Println(err)
		return resCourses.Courses, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resCourses.Courses, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resCourses.Courses, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resCourses.Courses, err
	}

	err = json.Unmarshal(body, &resCourses)
	if err != nil {
		fmt.Println(err)
		return resCourses.Courses, err
	}
	return resCourses.Courses, nil
}
