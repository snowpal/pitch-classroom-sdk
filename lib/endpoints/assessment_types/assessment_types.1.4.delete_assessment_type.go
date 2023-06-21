package assessmentTypes

import (
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
)

func DeleteAssessmentType(jwtToken string, assessmentTypeId string) error {
	route, err := helpers.GetRoute(lib.RouteAssessmentTypesDeleteAssessmentType, assessmentTypeId)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		return err
	}
	return nil
}
