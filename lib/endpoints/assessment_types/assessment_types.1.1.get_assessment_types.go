package assessmentTypes

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

func GetAssessmentTypes(jwtToken string, includeCounts bool) ([]response.AssessmentType, error) {
	resAssessmentTypes := response.AssessmentTypes{}
	route, err := helpers2.GetRoute(lib.RouteAssessmentTypesGetAssessmentTypes, strconv.FormatBool(includeCounts))
	if err != nil {
		fmt.Println(err)
		return resAssessmentTypes.AssessmentTypes, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAssessmentTypes.AssessmentTypes, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAssessmentTypes.AssessmentTypes, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAssessmentTypes.AssessmentTypes, err
	}

	err = json.Unmarshal(body, &resAssessmentTypes)
	if err != nil {
		fmt.Println(err)
		return resAssessmentTypes.AssessmentTypes, err
	}
	return resAssessmentTypes.AssessmentTypes, nil
}
