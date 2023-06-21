package teacherKeys

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetCourseGradesForStudents(
	jwtToken string,
	courseParam common.ResourceIdParam,
) (response.StudentGradeForCourseAndAssessment, error) {
	resStudentGradesForCourse := response.StudentGradeForCourseAndAssessment{}
	route, err := helpers2.GetRoute(
		lib.RouteTeacherKeysGetCourseGradesForStudents,
		courseParam.CourseId,
		courseParam.KeyId,
	)
	if err != nil {
		return resStudentGradesForCourse, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resStudentGradesForCourse, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resStudentGradesForCourse, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resStudentGradesForCourse, err
	}

	err = json.Unmarshal(body, &resStudentGradesForCourse)
	if err != nil {
		return resStudentGradesForCourse, err
	}
	return resStudentGradesForCourse, nil
}
