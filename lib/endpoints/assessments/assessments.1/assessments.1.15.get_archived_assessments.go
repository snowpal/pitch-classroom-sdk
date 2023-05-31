package assessments

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

func GetArchivedAssessments(jwtToken string, assessmentsParam request.GetAssessmentsParam) ([]response.Assessment, error) {
	resPods := response.Assessments{}
	route, err := helpers2.GetRoute(
		lib.RouteAssessmentsGetArchivedAssessments,
		strconv.Itoa(assessmentsParam.BatchIndex),
		assessmentsParam.KeyId,
		*assessmentsParam.CourseId,
	)
	if err != nil {
		fmt.Println(err)
		return resPods.Assessments, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resPods.Assessments, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	_, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPods.Assessments, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPods.Assessments, err
	}

	err = json.Unmarshal(body, &resPods)
	if err != nil {
		fmt.Println(err)
		return resPods.Assessments, err
	}
	return resPods.Assessments, nil
}
