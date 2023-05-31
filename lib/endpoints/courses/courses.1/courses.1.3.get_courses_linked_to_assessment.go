package courses

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetCoursesLinkedToAssessment(jwtToken string, blockParam common.ResourceIdParam) ([]response.Course, error) {
	resBlocks := response.Courses{}
	route, err := helpers2.GetRoute(
		lib.RouteCoursesGetCoursesLinkedToAssessment,
		blockParam.PodId,
		blockParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Courses, err
	}

	req, _ := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Courses, err
	}

	helpers2.AddUserHeaders(jwtToken, req)
	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Courses, err
	}

	defer helpers2.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
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
