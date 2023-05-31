package assessments

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetAssessments(jwtToken string, assessmentsParam request.GetAssessmentsParam) ([]response.Assessment, error) {
	resAssessments := response.Assessments{}
	route, err := helpers.GetRoute(
		lib.RouteAssessmentsGetAssessments,
		*assessmentsParam.CourseId,
		strconv.Itoa(assessmentsParam.BatchIndex),
		assessmentsParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resAssessments.Assessments, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAssessments.Assessments, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAssessments.Assessments, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAssessments.Assessments, err
	}

	err = json.Unmarshal(body, &resAssessments)
	if err != nil {
		fmt.Println(err)
		return resAssessments.Assessments, err
	}
	return resAssessments.Assessments, nil
}
