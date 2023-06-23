package setupnewuser

import (
	"fmt"

	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/assessments/assessments.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/courses/courses.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

var KeyName = "Math Class"
var CourseName = "Trigonometry"
var AssessmentName = "Trigonometric Functions"

func CreateData(user response.User) (KeyWithCourses, error) {
	var KeyWithCourses KeyWithCourses

	log.Info("Creating teacher key")
	key, err := keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: KeyName,
			Type: lib.TeacherKeyType,
		})
	if err != nil {
		return KeyWithCourses, err
	}

	var CourseWithAssessments CourseWithAssessments

	log.Info(fmt.Sprintf("Creating course inside %s key.", KeyName))
	var course response.Course
	course, err = courses.AddCourse(
		user.JwtToken,
		request.AddCourseReqBody{Name: CourseName},
		key.ID)
	if err != nil {
		return KeyWithCourses, err
	}
	CourseWithAssessments.Course = course

	log.Info(fmt.Sprintf("Creating assessment inside %s course.", CourseName))
	var assessment response.Assessment
	assessment, err = assessments.AddAssessment(
		user.JwtToken,
		request.AddAssessmentReqBody{Name: AssessmentName},
		common.ResourceIdParam{CourseId: course.ID, KeyId: course.Key.ID},
	)
	if err != nil {
		return KeyWithCourses, err
	}
	CourseWithAssessments.Assessments = append(CourseWithAssessments.Assessments, assessment)

	KeyWithCourses.Courses = append(KeyWithCourses.Courses, CourseWithAssessments)

	return KeyWithCourses, err
}
