package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

var KeyName = "Math Class"
var CourseName = "Trigonometry"
var AssessmentName = "Trigonometric Functions"

func CreateContent(user response.User) (KeyWithCourses, error) {
	var err error
	var KeyWithCourses KeyWithCourses

	log.Info("Creating teacher key")
	key, err := recipes.AddTeacherKey(user, KeyName)
	if err != nil {
		return KeyWithCourses, err
	}

	var CourseWithAssessments CourseWithAssessments

	log.Info("Creating course inside ", KeyName, " key.")
	var course response.Course
	course, err = recipes.AddCourse(user, CourseName, key)
	if err != nil {
		return KeyWithCourses, err
	}
	CourseWithAssessments.Course = course

	log.Info("Creating assessment inside ", CourseName, " course.")
	var assessment response.Assessment
	assessment, err = recipes.AddAssessmentToCourse(user, AssessmentName, course)
	if err != nil {
		return KeyWithCourses, err
	}
	CourseWithAssessments.Assessments = append(CourseWithAssessments.Assessments, assessment)

	KeyWithCourses.Courses = append(KeyWithCourses.Courses, CourseWithAssessments)

	return KeyWithCourses, err
}
