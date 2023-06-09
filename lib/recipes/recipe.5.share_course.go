package recipes

import (
	"fmt"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/collaboration/collaboration.1.courses"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/courses/courses.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/notifications"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
	keys "github.com/snowpal/pitch-classroom-sdk/lib/endpoints/keys/keys.1"
	user2 "github.com/snowpal/pitch-classroom-sdk/lib/endpoints/user"
)

const (
	KeyName           = "Math Class"
	CourseName        = "Algebra"
	UpdatedCourseName = "Geometry"
)

func AddStudentAndTeacher() {
	log.Info("Objective: Create course, add student & teacher to that course.")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	var user response.User
	user, err = recipes.SignIn(lib.ApiUser1, lib.Password)
	if err != nil {
		return
	}

	log.Info("Add student & teacher to a course")
	recipes.SleepBefore()
	var course response.Course
	course, err = addStudentAndTeacher(user)
	if err != nil {
		return
	}

	var student response.User
	student, err = getStudent(user, course)
	if err != nil {
		return
	}

	log.Info("Show notifications as student")
	recipes.SleepBefore()
	err = showNotificationsAsStudent(student)
	if err != nil {
		return
	}
	log.Info(".Notifications for the recent share displayed successfully")
	recipes.SleepAfter()
}

func getStudent(user response.User, course response.Course) (response.User, error) {
	var student response.User
	resCourse, err := collaboration.GetCourseCollaborators(
		user.JwtToken,
		common.ResourceIdParam{
			CourseId: course.ID,
			KeyId:    course.Key.ID,
		})
	if err != nil {
		return student, err
	}
	allUsers, err := user2.GetUsers(user.JwtToken)
	for _, sharedUser := range *resCourse.SharedUsers {
		for _, userInAll := range allUsers {
			if sharedUser.ID == userInAll.ID && sharedUser.Acl == lib.StudentAcl {
				student = userInAll
				break
			}
		}
	}
	if err != nil {
		return student, err
	}

	student, err = recipes.SignIn(student.Email, lib.Password)
	if err != nil {
		return student, err
	}
	return student, nil
}

func addStudentAndTeacher(user response.User) (response.Course, error) {
	var course response.Course
	key, err := keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: KeyName,
			Type: lib.TeacherKeyType,
		})
	if err != nil {
		return course, err
	}
	course, err = courses.AddCourse(
		user.JwtToken,
		request.AddCourseReqBody{Name: CourseName},
		key.ID)
	if err != nil {
		return course, err
	}
	err = recipes.SearchUserAndShareCourse(user, course, lib.ApiUser2, lib.StudentAcl)
	if err != nil {
		return course, err
	}
	err = recipes.SearchUserAndShareCourse(user, course, lib.ApiUser3, lib.TeacherAcl)
	if err != nil {
		return course, err
	}
	return course, nil
}

func showNotificationsAsStudent(student response.User) error {
	unreadNotifications, err := notifications.GetNotifications(student.JwtToken)
	if err != nil {
		return err
	}
	for index, notification := range unreadNotifications {
		if notification.Type == "acl" {
			log.Info(fmt.Sprintf(".Notification %d: %s", index, notification.Text))
		}
	}
	return nil
}
