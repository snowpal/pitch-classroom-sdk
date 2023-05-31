package recipes

import (
	"fmt"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/collaboration/collaboration.1.courses"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/courses/courses.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/notifications"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
	user2 "github.com/snowpal/pitch-classroom-sdk/lib/endpoints/user"
)

const (
	KeyName           = "Diwali Festival"
	CourseName        = "Diwali Function"
	UpdatedCourseName = "Diwali Celebration"
)

func ShareCourse() {
	log.Info("Objective: Create course, share users as read & write, make 1 of them as admin.")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	log.Info("Share a course")
	recipes.SleepBefore()
	var course response.Course
	course, err = shareCourse(user)
	if err != nil {
		return
	}

	writeUser, err := getWriteUser(user, course)
	fmt.Println(user.JwtToken)
	if err != nil {
		return
	}

	log.Info("Show notifications as write user")
	recipes.SleepBefore()
	err = showNotificationsAsWriteUser(writeUser)
	if err != nil {
		return
	}
	log.Printf(".Notifications for the recent share displayed successfully")
	recipes.SleepAfter()

	log.Printf("Update course name as a write user")
	recipes.SleepBefore()
	var resCourse response.Course
	resCourse, err = updateCourseAsWriteUser(writeUser, course)
	if err != nil {
		return
	}
	log.Printf(".Write user updated course name to %s successfully", resCourse.Name)
	recipes.SleepAfter()

	log.Printf("Grant admin access to a user with read access")
	err = makeReadUserAsAdmin(user, course)
	if err != nil {
		return
	}
	log.Printf(".Admin access has been granted successfully")
}

func getWriteUser(user response.User, course response.Course) (response.User, error) {
	var writeUser response.User
	resCourse, err := collaboration.GetCourseCollaborators(
		user.JwtToken,
		common.ResourceIdParam{
			CourseId: course.ID,
			KeyId:    course.Key.ID,
		})
	if err != nil {
		return writeUser, err
	}
	allUsers, err := user2.GetUsers(user.JwtToken)
	for _, sharedUser := range *resCourse.SharedUsers {
		for _, userInAll := range allUsers {
			if sharedUser.ID == userInAll.ID {
				writeUser = userInAll
				break
			}
		}
	}
	if err != nil {
		return writeUser, err
	}

	writeUser, err = recipes.SignIn(writeUser.Email, lib.Password)
	if err != nil {
		return writeUser, err
	}
	return writeUser, nil
}

func shareCourse(user response.User) (response.Course, error) {
	var course response.Course
	key, err := recipes.AddTeacherKey(user, KeyName)
	if err != nil {
		return course, err
	}
	course, err = recipes.AddCourse(user, CourseName, key)
	if err != nil {
		return course, err
	}
	err = recipes.SearchUserAndShareCourse(user, course, "api_read_user", lib.ReadAcl)
	if err != nil {
		return course, err
	}
	err = recipes.SearchUserAndShareCourse(user, course, "api_write_user", lib.WriteAcl)
	if err != nil {
		return course, err
	}
	return course, nil
}

func showNotificationsAsWriteUser(writeUser response.User) error {
	unreadNotifications, err := notifications.GetNotifications(writeUser.JwtToken)
	if err != nil {
		return err
	}
	fmt.Println(len(unreadNotifications))
	for index, notification := range unreadNotifications {
		if notification.Type == "acl" {
			log.Printf(".Notification %d: %s", index, notification.Text)
		}
	}
	return nil
}

func updateCourseAsWriteUser(writeUser response.User, course response.Course) (response.Course, error) {
	const (
		SystemKeyType       = "system"
		customSystemKeyType = "SharedCustomKey"
	)
	systemKeys, _ := keys.GetKeysFilteredByType(writeUser.JwtToken, SystemKeyType)
	var customSystemKey response.Key
	for _, systemKey := range systemKeys {
		if systemKey.Type == customSystemKeyType {
			customSystemKey = systemKey
			break
		}
	}
	updatedCourseName := UpdatedCourseName
	resCourse, err := courses.UpdateCourse(
		writeUser.JwtToken,
		courses.UpdateCourseReqBody{Name: &updatedCourseName},
		common.ResourceIdParam{
			CourseId: course.ID,
			KeyId:    customSystemKey.ID,
		})
	if err != nil {
		return resCourse, err
	}
	return resCourse, nil
}

func makeReadUserAsAdmin(user response.User, course response.Course) error {
	resCourse, err := collaboration.GetCourseCollaborators(
		user.JwtToken,
		common.ResourceIdParam{
			CourseId: course.ID,
			KeyId:    course.Key.ID,
		})
	if err != nil {
		return err
	}

	var readUser response.SharedUser
	for _, sharedUser := range *resCourse.SharedUsers {
		if sharedUser.Acl == lib.ReadAcl {
			readUser = sharedUser
			break
		}
	}

	_, err = collaboration.UpdateCourseAcl(
		user.JwtToken,
		request.CourseAclReqBody{Acl: lib.AdminAcl},
		common.AclParam{
			UserId: readUser.ID,
			ResourceIds: common.ResourceIdParam{
				CourseId: course.ID,
				KeyId:    course.Key.ID,
			},
		})
	if err != nil {
		return err
	}
	return nil
}
