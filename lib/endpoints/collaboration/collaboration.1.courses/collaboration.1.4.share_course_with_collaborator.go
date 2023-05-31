package collaboration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func ShareCourseWithCollaborator(
	jwtToken string,
	reqBody request.CourseAclReqBody,
	courseAclParam common.AclParam,
) (response.Course, error) {
	resCourse := response.Course{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resCourse, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteCollaborationShareCourseWithCollaborator,
		courseAclParam.ResourceIds.CourseId,
		courseAclParam.UserId,
		courseAclParam.ResourceIds.KeyId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resCourse, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resCourse, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resCourse, err
	}

	err = json.Unmarshal(body, &resCourse)
	if err != nil {
		fmt.Println(err)
		return resCourse, err
	}
	return resCourse, nil
}
