package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/assessments/assessments.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/courses/courses.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"

	recipes "github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
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

	user, err := recipes.SignIn(lib.ApiUser1, lib.Password)
	if err != nil {
		return
	}

	var newKey response.Key
	log.Info("Add a new teacher key")
	recipes.SleepBefore()
	newKey, err = recipes.AddTeacherKey(user, Key1Name)
	if err != nil {
		return
	}
	log.Printf(".Key, %s is added successfully.", newKey.Name)
	recipes.SleepAfter()

	var (
		newCourse     response.Course
		newAssessment response.Assessment
	)
	newCourse, newAssessment, err = addCoursesAndAssessments(user, newKey)
	if err != nil {
		return
	}

	log.Info("Add another key")
	recipes.SleepBefore()
	var anotherKey response.Key
	anotherKey, err = recipes.AddStudentKey(user, AnotherKeyName)
	if err != nil {
		return
	}

	log.Info("Add course")
	recipes.SleepBefore()
	var anotherCourse response.Course
	anotherCourse, err = recipes.AddCourse(user, AnotherCourseName, newKey)
	if err != nil {
		return
	}
	log.Printf(".Course, %s is created successfully.", newCourse.Name)
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
	log.Printf(".Course, %s is linked successfully to %s Key.", newCourse.Name, anotherKey.Name)
	recipes.SleepAfter()
	return nil
}

func addCoursesAndAssessments(user response.User, newKey response.Key) (response.Course, response.Assessment, error) {
	var (
		assessment response.Assessment
		course     response.Course
	)
	log.Info("Add a new course")
	recipes.SleepBefore()
	newCourse, err := recipes.AddCourse(user, Course1Name, newKey)
	if err != nil {
		return course, assessment, err
	}
	log.Printf(".Course, %s is created successfully.", newCourse.Name)
	recipes.SleepAfter()

	log.Info("Add a new course assessment in this course")
	recipes.SleepBefore()
	var newAssessment response.Assessment
	newAssessment, err = assessments.AddAssessment(user.JwtToken,
		request.AddAssessmentReqBody{
			Name: Assessment1Name,
		},
		common.ResourceIdParam{
			CourseId: newCourse.ID,
			KeyId:    newKey.ID,
		})
	if err != nil {
		return course, assessment, err
	}
	log.Printf(".Assessment, %s is created successfully in %s Course.", newAssessment.Name, newCourse.Name)
	recipes.SleepAfter()
	return newCourse, newAssessment, nil
}
