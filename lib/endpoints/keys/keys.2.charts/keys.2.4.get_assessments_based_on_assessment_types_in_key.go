package keys

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetAssessmentsBasedOnAssessmentTypesInKey(jwtToken string, keyId string) (response.AssessmentTypesKey, error) {
	resAssessmentTypesKey := response.AssessmentTypesKey{}
	route, err := helpers2.GetRoute(lib.RouteKeysGetAssessmentsBasedOnAssessmentTypesInKey, keyId)
	if err != nil {
		return resAssessmentTypesKey, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resAssessmentTypesKey, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resAssessmentTypesKey, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resAssessmentTypesKey, err
	}

	err = json.Unmarshal(body, &resAssessmentTypesKey)
	if err != nil {
		return resAssessmentTypesKey, err
	}
	return resAssessmentTypesKey, nil
}
