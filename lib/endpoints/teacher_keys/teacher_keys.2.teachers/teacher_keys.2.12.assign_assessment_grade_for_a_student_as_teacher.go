package teacherKeys

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func AssignAssessmentGradeForAStudentAsTeacher(
	jwtToken string,
	reqBody request.UpdateGradeReqBody,
	classroomParam request.ClassroomIdParam,
) (response.UpdateAssessmentGrade, error) {
	resAssessmentGrade := response.UpdateAssessmentGrade{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		return resAssessmentGrade, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		lib.RouteTeacherKeysAssignAssessmentGradeForAStudentAsTeacher,
		classroomParam.ResourceIds.AssessmentId,
		classroomParam.StudentId,
		classroomParam.ResourceIds.KeyId,
		classroomParam.ResourceIds.CourseId,
	)
	if err != nil {
		return resAssessmentGrade, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		return resAssessmentGrade, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		return resAssessmentGrade, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resAssessmentGrade, err
	}

	err = json.Unmarshal(body, &resAssessmentGrade)
	if err != nil {
		return resAssessmentGrade, err
	}
	return resAssessmentGrade, nil
}
