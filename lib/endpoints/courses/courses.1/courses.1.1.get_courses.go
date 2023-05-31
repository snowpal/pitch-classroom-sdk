package courses

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetCourses(jwtToken string, blockParam request.GetCoursesParam) ([]response.Course, error) {
	resBlocks := response.Courses{}
	route, err := helpers2.GetRoute(
		lib.RouteCoursesGetCourses,
		blockParam.KeyId,
		blockParam.Filter,
		strconv.Itoa(blockParam.BatchIndex),
		strconv.FormatBool(blockParam.WriteOrHigherAcl),
	)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Courses, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Courses, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Courses, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Courses, err
	}

	err = json.Unmarshal(body, &resBlocks)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Courses, err
	}
	return resBlocks.Courses, nil
}
