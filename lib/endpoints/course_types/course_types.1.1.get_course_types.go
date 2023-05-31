package block_types

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetCourseTypes(jwtToken string, includeCounts bool) ([]response.CourseType, error) {
	resBlockTypes := response.CourseTypes{}
	route, err := helpers2.GetRoute(lib.RouteCourseTypesGetCourseTypes, strconv.FormatBool(includeCounts))
	if err != nil {
		fmt.Println(err)
		return resBlockTypes.CourseTypes, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlockTypes.CourseTypes, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockTypes.CourseTypes, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlockTypes.CourseTypes, err
	}

	err = json.Unmarshal(body, &resBlockTypes)
	if err != nil {
		fmt.Println(err)
		return resBlockTypes.CourseTypes, err
	}
	return resBlockTypes.CourseTypes, nil
}
