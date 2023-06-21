package teacherKeys

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetCourseAndAssessmentsGradesForAStudentAsTeacher(
	jwtToken string,
	classroomParam request.ClassroomIdParam,
) (response.StudentGradeForCourseAndAssessment, error) {
	resStudentGrades := response.StudentGradeForCourseAndAssessment{}
	route, err := helpers2.GetRoute(
		lib.RouteTeacherKeysGetCourseAndAssessmentsGradesForAStudentAsTeacher,
		classroomParam.ResourceIds.CourseId,
		classroomParam.StudentId,
		classroomParam.ResourceIds.KeyId,
	)
	if err != nil {
		return resStudentGrades, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resStudentGrades, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resStudentGrades, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resStudentGrades, err
	}

	err = json.Unmarshal(body, &resStudentGrades)
	if err != nil {
		return resStudentGrades, err
	}
	return resStudentGrades, nil
}
