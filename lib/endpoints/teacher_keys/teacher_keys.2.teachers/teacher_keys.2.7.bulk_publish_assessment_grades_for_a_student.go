package teacherKeys

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
)

type PublishStudentGradeReqBody struct {
	PodIds string `json:"podIds"`
}

func BulkPublishAssessmentGradesForAStudent(
	jwtToken string,
	reqBody PublishStudentGradeReqBody,
	classroomParam request.ClassroomIdParam,
) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		lib.RouteTeacherKeysBulkPublishAssessmentGradesForAStudent,
		classroomParam.StudentId,
		classroomParam.ResourceIds.KeyId,
		classroomParam.ResourceIds.CourseId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
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
