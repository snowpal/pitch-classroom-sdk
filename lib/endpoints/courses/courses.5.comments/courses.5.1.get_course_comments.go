package courses

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func GetCourseComments(jwtToken string, commentParam request.CommentIdParam) ([]response.Comment, error) {
	resComments := response.Comments{}
	route, err := helpers2.GetRoute(
		lib.RouteCoursesGetCourseComments,
		*commentParam.CourseId,
		commentParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resComments.Comments, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resComments.Comments, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resComments.Comments, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resComments.Comments, err
	}

	err = json.Unmarshal(body, &resComments)
	if err != nil {
		fmt.Println(err)
		return resComments.Comments, err
	}
	return resComments.Comments, nil
}
