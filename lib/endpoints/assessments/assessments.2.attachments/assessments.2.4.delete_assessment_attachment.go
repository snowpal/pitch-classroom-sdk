package assessments

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
)

func DeleteAssessmentAttachment(jwtToken string, attachmentParam request.AttachmentParam) error {
	route, err := helpers.GetRoute(
		lib.RouteAssessmentsDeleteAssessmentAttachment,
		*attachmentParam.AttachmentId,
		attachmentParam.KeyId,
		*attachmentParam.BlockId,
		*attachmentParam.PodId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
