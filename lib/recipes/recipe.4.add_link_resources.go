package recipes

import (
	"fmt"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/assessments/assessments.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/courses/courses.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"

	log "github.com/sirupsen/logrus"
	keys "github.com/snowpal/pitch-classroom-sdk/lib/endpoints/keys/keys.1"
	recipes "github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

const (
	Key1Name          = "Taxes"
	AnotherKeyName    = "State Taxes"
	Course1Name       = "Form 1040"
	AnotherCourseName = "Form 1120S"
	Assessment1Name   = "Expenses"
)

// AddAndLinkResources Add course & assessment to a key and link them into another key
func AddAndLinkResources() {
	log.Info("Objective: Add keys and courses, and link courses")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	var user response.User
	user, err = recipes.SignIn(lib.ApiUser1, lib.Password)
	if err != nil {
		return
	}

	var key response.Key
	log.Info("Add a new teacher key")
	recipes.SleepBefore()
	key, err = keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: Key1Name,
			Type: lib.TeacherKeyType,
		})
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Key, %s is added successfully.", key.Name))
	recipes.SleepAfter()

	var (
		newCourse     response.Course
		newAssessment response.Assessment
	)
	newCourse, newAssessment, err = addCoursesAndAssessments(user, key)
	if err != nil {
		return
	}

	log.Info("Add another key")
	recipes.SleepBefore()
	var anotherKey response.Key
	anotherKey, err = keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: AnotherKeyName,
			Type: lib.TeacherKeyType,
		})
	if err != nil {
		return
	}

	log.Info("Add course")
	recipes.SleepBefore()
	var anotherCourse response.Course
	anotherCourse, err = courses.AddCourse(
		user.JwtToken,
		request.AddCourseReqBody{Name: AnotherCourseName},
		key.ID)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Course, %s is created successfully.", newCourse.Name))
	recipes.SleepAfter()

	err = linkResources(user, anotherKey, anotherCourse, newCourse, newAssessment)
	if err != nil {
		return
	}
}

func linkResources(
	user response.User,
	anotherKey response.Key,
	anotherCourse response.Course,
	newCourse response.Course,
	newAssessment response.Assessment,
) error {
	log.Info("Link course into the other key")
	recipes.SleepBefore()
	err := courses.LinkCourseToKey(user.JwtToken,
		common.ResourceIdParam{
			CourseId: newCourse.ID,
			KeyId:    anotherKey.ID,
		})
	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf(".Course, %s is linked successfully to %s Key.", newCourse.Name, anotherKey.Name))
	recipes.SleepAfter()
	return nil
}

func addCoursesAndAssessments(user response.User, key response.Key) (response.Course, response.Assessment, error) {
	var (
		assessment response.Assessment
		course     response.Course
		err        error
	)
	log.Info("Add a new course")
	recipes.SleepBefore()
	course, err = courses.AddCourse(
		user.JwtToken,
		request.AddCourseReqBody{Name: Course1Name},
		key.ID)
	if err != nil {
		return course, assessment, err
	}
	log.Info(fmt.Sprintf(".Course, %s is created successfully.", course.Name))
	recipes.SleepAfter()

	log.Info("Add a new course assessment in this course")
	recipes.SleepBefore()
	assessment, err = assessments.AddAssessment(user.JwtToken,
		request.AddAssessmentReqBody{Name: Assessment1Name},
		common.ResourceIdParam{
			CourseId: course.ID,
			KeyId:    key.ID,
		})
	if err != nil {
		return course, assessment, err
	}
	log.Info(fmt.Sprintf(".Assessment, %s is created successfully in %s Course.", assessment.Name, course.Name))
	recipes.SleepAfter()
	return course, assessment, nil
}
