package relations

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
)

func unrelateAssessmentFromAssessment(jwtToken string, route string) error {
	req, err := http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UnrelateAssessmentFromAssessment(jwtToken string, relationParam request.AssessmentToAssessmentRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsUnrelateAssessmentFromAssessment,
		relationParam.PodId,
		relationParam.SourceKeyId,
		relationParam.SourceBlockId,
		relationParam.TargetPodId,
		relationParam.TargetKeyId,
		relationParam.TargetBlockId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = unrelateAssessmentFromAssessment(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
