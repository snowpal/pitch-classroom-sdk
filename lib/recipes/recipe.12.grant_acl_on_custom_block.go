package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/courses/courses.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"

	recipes "github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

const (
	CopyKeyName    = "Insurance"
	CopyCourseName = "Car Insurance"
)

func GrantAclOnCustomCourse() {
	log.Info("Objective: Add Custom Course, Share Course, Grant Read Access, Copy Course, Grant Admin Access")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	key, err := recipes.AddTeacherKey(user, CopyKeyName)
	if err != nil {
		return
	}

	log.Info("Add custom course")
	recipes.SleepBefore()
	course, err := recipes.AddCourse(user, CopyCourseName, key)
	if err != nil {
		return
	}
	log.Printf(".Course %s added successfully", course.Name)
	recipes.SleepAfter()

	log.Info("Share course with read access")
	recipes.SleepBefore()
	err = recipes.SearchUserAndShareCourse(user, course, "api_read_user", lib.ReadAcl)
	if err != nil {
		return
	}
	log.Printf(".Course %s shared with %s with read access level", course.Name, lib.ReadUser)
	recipes.SleepAfter()

	log.Info("Copy course and see acl is not copied")
	recipes.SleepBefore()
	anotherCourse, err := copyCourse(user, course)
	if err != nil {
		return
	}
	log.Printf(".Course %s copied but %s don't have access on copied course", course.Name, lib.ReadUser)
	recipes.SleepAfter()

	log.Info("Share course with admin access")
	recipes.SleepBefore()
	err = recipes.SearchUserAndShareCourse(user, anotherCourse, "api_admin_user", lib.AdminAcl)
	if err != nil {
		return
	}
	log.Printf(".Course %s shared with %s with admin access", course.Name, lib.ReadUser)
	recipes.SleepAfter()
}

func copyCourse(user response.User, course response.Course) (response.Course, error) {
	resCourse, err := courses.CopyCourse(
		user.JwtToken,
		request.CopyMoveCourseParam{
			CourseId:       course.ID,
			KeyId:          course.Key.ID,
			TargetKeyId:    course.Key.ID,
			AllAssessments: true,
			AllTasks:       true,
			AllChecklists:  true,
		},
	)
	if err != nil {
		return resCourse, err
	}
	return resCourse, nil
}
