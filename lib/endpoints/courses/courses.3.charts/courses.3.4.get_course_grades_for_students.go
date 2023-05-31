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

type CourseGradesForStudents struct {
	Course      response.CourseGrade       `json:"course"`
	Assessments []response.AssessmentGrade `json:"assessments"`
}

func GetCourseGradesForStudents(jwtToken string, gradeParam common.ResourceIdParam) (CourseGradesForStudents, error) {
	resCourseGrades := CourseGradesForStudents{}
	route, err := helpers2.GetRoute(
		lib.RouteCoursesGetCourseGradesForStudents,
		gradeParam.CourseId,
		gradeParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resCourseGrades, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resCourseGrades, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resCourseGrades, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resCourseGrades, err
	}

	err = json.Unmarshal(body, &resCourseGrades)
	if err != nil {
		fmt.Println(err)
		return resCourseGrades, err
	}
	return resCourseGrades, nil
}
