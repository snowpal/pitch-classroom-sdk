package assessments

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	common2 "github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

type AssessmentGradesForStudents struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"courseName"`
	Key        common2.SlimKey        `json:"key"`
	Assessment common2.SlimAssessment `json:"assessment"`
	Students   []response.Student     `json:"students"`
}

func GetAssessmentGradesForStudents(jwtToken string, assessmentParam common2.ResourceIdParam) (AssessmentGradesForStudents, error) {
	resAssessmentGrades := AssessmentGradesForStudents{}
	route, err := helpers2.GetRoute(
		lib.RouteAssessmentsGetAssessmentGradesForStudents,
		assessmentParam.AssessmentId,
		assessmentParam.KeyId,
		assessmentParam.CourseId,
	)
	if err != nil {
		return resAssessmentGrades, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resAssessmentGrades, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resAssessmentGrades, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resAssessmentGrades, err
	}

	err = json.Unmarshal(body, &resAssessmentGrades)
	if err != nil {
		return resAssessmentGrades, err
	}
	return resAssessmentGrades, nil
}
