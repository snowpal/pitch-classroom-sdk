package assessments

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	helpers2 "github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func UpdateAssessmentComment(
	jwtToken string,
	reqBody request.CommentReqBody,
	commentParam request.CommentIdParam,
) (response.Comment, error) {
	resComment := response.Comment{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resComment, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteAssessmentsUpdateAssessmentComment,
		*commentParam.CommentId,
		commentParam.KeyId,
		*commentParam.CourseId,
		*commentParam.AssessmentId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resComment, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resComment, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resComment, err
	}

	err = json.Unmarshal(body, &resComment)
	if err != nil {
		fmt.Println(err)
		return resComment, err
	}
	return resComment, nil
}
