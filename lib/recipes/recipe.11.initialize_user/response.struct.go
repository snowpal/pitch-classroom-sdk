package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"
)

type KeyWithCourses struct {
	Key     response.Key
	Courses []CourseWithAssessments
}

type CourseWithAssessments struct {
	Course      response.Course
	Assessments []response.Assessment
}
