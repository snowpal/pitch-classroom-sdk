package courses

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

type BlockGradesForStudents struct {
	Course      response.CourseGrade       `json:"course"`
	Assessments []response.AssessmentGrade `json:"pods"`
}

func GetCourseGradesForStudents(jwtToken string, gradeParam common.ResourceIdParam) (BlockGradesForStudents, error) {
	resBlockGrades := BlockGradesForStudents{}
	route, err := helpers2.GetRoute(
		lib.RouteCoursesGetCourseGradesForStudents,
		gradeParam.CourseId,
		gradeParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resBlockGrades, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlockGrades, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockGrades, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlockGrades, err
	}

	err = json.Unmarshal(body, &resBlockGrades)
	if err != nil {
		fmt.Println(err)
		return resBlockGrades, err
	}
	return resBlockGrades, nil
}
