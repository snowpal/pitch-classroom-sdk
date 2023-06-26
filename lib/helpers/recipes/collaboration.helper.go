package recipes

import (
	"strings"

	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/collaboration/collaboration.1.courses"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func SearchUserAndShareCourse(user response.User, course response.Course, email string, acl string) error {
	courseIdParam := common.ResourceIdParam{
		CourseId: course.ID,
		KeyId:    course.Key.ID,
	}

	searchedUsers, err := collaboration.GetUsersThisCourseCanBeSharedWith(
		user.JwtToken,
		common.SearchUsersParam{
			SearchToken: strings.Split(email, "@")[0],
			ResourceIds: courseIdParam,
		})
	if err != nil {
		return err
	}

	// For the purpose of this recipe, it does not matter which user from the list we end up picking, hence we go with
	// the first one.
	_, err = collaboration.ShareCourseWithCollaborator(
		user.JwtToken,
		request.CourseAclReqBody{Acl: acl},
		common.AclParam{
			UserId:      searchedUsers[0].ID,
			ResourceIds: courseIdParam,
		})
	if err != nil {
		return err
	}
	return nil
}
