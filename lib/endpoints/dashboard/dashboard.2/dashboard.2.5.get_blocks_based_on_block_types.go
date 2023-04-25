package dashboard

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetBlocksBasedOnBlockTypes(jwtToken string) ([]response.BlockTypesKey, error) {
	resBlockTypesKeys := response.BlockTypesKeys{}
	route, err := helpers2.GetRoute(lib.RouteDashboardGetCoursesBasedOnCourseTypes)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKeys.Keys, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKeys.Keys, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKeys.Keys, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKeys.Keys, err
	}

	err = json.Unmarshal(body, &resBlockTypesKeys)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKeys.Keys, err
	}
	return resBlockTypesKeys.Keys, nil
}
