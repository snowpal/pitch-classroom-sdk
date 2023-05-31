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

func GetCourseGradesForStudents(
	jwtToken string,
	blockParam common.ResourceIdParam,
) (response.StudentGradeForCourseAndAssessment, error) {
	resStudentGradesForBlock := response.StudentGradeForCourseAndAssessment{}
	route, err := helpers2.GetRoute(
		lib.RouteTeacherKeysGetCourseGradesForStudents,
		blockParam.BlockId,
		blockParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForBlock, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForBlock, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForBlock, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForBlock, err
	}

	err = json.Unmarshal(body, &resStudentGradesForBlock)
	if err != nil {
		fmt.Println(err)
		return resStudentGradesForBlock, err
	}
	return resStudentGradesForBlock, nil
}
