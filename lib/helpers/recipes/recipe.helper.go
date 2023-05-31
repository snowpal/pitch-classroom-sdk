package recipes

import (
	"fmt"
	"time"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/assessments/assessments.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/collaboration/collaboration.1.courses"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/courses/courses.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

func sleepWindow(sleepTime time.Duration) {
	time.Sleep(time.Second * sleepTime)
}

func SleepBefore() {
	sleepWindow(1)
}

func SleepAfter() {
	sleepWindow(2)
}

// ValidateDependencies We require that the first recipe be run before anything else as it registers a bunch of users.
// To verify if it was actually run, we do this "random" check.
func ValidateDependencies() (response.User, error) {
	user, err := SignIn(lib.DefaultEmail, lib.Password)
	fmt.Println(user)
	fmt.Println(err)
	if err != nil {
		return user, err
	}

	return user, nil
}

func addKey(user response.User, keyName string, keyType string) (response.Key, error) {
	newKey, err := keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: keyName,
			Type: keyType,
		})
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddStudentKey(user response.User, keyName string) (response.Key, error) {
	newKey, err := addKey(user, keyName, lib.StudentKeyType)
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddTeacherKey(user response.User, keyName string) (response.Key, error) {
	newKey, err := addKey(user, keyName, lib.TeacherKeyType)
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddCourse(user response.User, coursesName string, key response.Key) (response.Course, error) {
	newCourse, err := courses.AddCourse(
		user.JwtToken,
		request.AddCourseReqBody{Name: coursesName},
		key.ID)
	if err != nil {
		return newCourse, err
	}
	return newCourse, nil
}

func AddAssessmentToCourse(user response.User, assessmentName string, course response.Course) (response.Assessment, error) {
	newAssessment, err := assessments.AddAssessment(
		user.JwtToken,
		request.AddAssessmentReqBody{Name: assessmentName},
		common.ResourceIdParam{CourseId: course.ID, KeyId: course.Key.ID})
	if err != nil {
		return newAssessment, err
	}
	return newAssessment, nil
}

func SearchUserAndShareCourse(user response.User, course response.Course, searchToken string, acl string) error {
	courseIdParam := common.ResourceIdParam{
		CourseId: course.ID,
		KeyId:    course.Key.ID,
	}

	searchedUsers, err := collaboration.GetUsersThisCourseCanBeSharedWith(
		user.JwtToken,
		common.SearchUsersParam{
			SearchToken: searchToken,
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
