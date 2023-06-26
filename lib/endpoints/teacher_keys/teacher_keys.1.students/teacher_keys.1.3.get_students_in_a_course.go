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

func GetStudentsInACourse(jwtToken string, courseParam common.ResourceIdParam) ([]response.Student, error) {
	resStudents := response.Students{}
	route, err := helpers2.GetRoute(lib.RouteTeacherKeysGetStudentsInACourse, courseParam.CourseId, courseParam.KeyId)
	if err != nil {
		return resStudents.Students, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resStudents.Students, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resStudents.Students, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resStudents.Students, err
	}

	err = json.Unmarshal(body, &resStudents)
	if err != nil {
		return resStudents.Students, err
	}
	return resStudents.Students, nil
}
