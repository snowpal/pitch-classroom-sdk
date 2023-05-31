package teacherKeys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetAssessmentGradesForStudents(
	jwtToken string,
	assessmentParam common.ResourceIdParam,
) (response.StudentGradeForCourseAndAssessment, error) {
	resStudentGradesForPod := response.StudentGradeForCourseAndAssessment{}
	route, err := helpers2.GetRoute(
		lib.RouteTeacherKeysGetAssessmentGradesForStudents,
		assessmentParam.AssessmentId,
		assessmentParam.KeyId,
		assessmentParam.CourseId,
	)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}

	err = json.Unmarshal(body, &resStudentGradesForPod)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForPod, err
	}
	return resStudentGradesForPod, nil
}
