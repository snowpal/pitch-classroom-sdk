package collaboration

import (
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

type CourseBulkShareReqBody struct {
	Acl       string `json:"courseAcl"`
	CourseIds string `json:"courseIds"`
}

func ShareCoursesWithCollaborators(
	jwtToken string,
	reqBody CourseBulkShareReqBody,
	courseAclParam common.AclParam,
) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		return err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		lib.RouteCollaborationBulkShareCoursesWithCollaborators,
		courseAclParam.UserId,
		courseAclParam.ResourceIds.KeyId,
	)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
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
