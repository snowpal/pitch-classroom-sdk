package relations

import (
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
)

func RelateCourseToKey(jwtToken string, relationParam request.KeyToCourseRelationParam) error {
	route, err := helpers.GetRoute(
		lib.RouteRelationsRelateCourseToKey,
		relationParam.KeyId,
		relationParam.TargetCourseId,
	)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, nil)
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
