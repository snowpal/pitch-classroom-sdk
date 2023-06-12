package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/assessments/assessments.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	gradingSystems "github.com/snowpal/pitch-classroom-sdk/lib/endpoints/grading_systems"
	teacherKeys "github.com/snowpal/pitch-classroom-sdk/lib/endpoints/teacher_keys/teacher_keys.2.teachers"

	log "github.com/sirupsen/logrus"
)

const (
	TeacherKeyName        = "School"
	TeacherCourseName     = "Math"
	TeacherAssessmentName = "Final Exam"
)

func PublishStudentGrade() {
	log.Info("Objective: Assign and Publish Assessment Grades for a Student")
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
	key, err = recipes.AddTeacherKey(user, TeacherKeyName)
	if err != nil {
		return
	}

	var course response.Course
	course, err = recipes.AddCourse(user, TeacherCourseName, key)
	if err != nil {
		return
	}

	var assessment response.Assessment
	assessment, err = recipes.AddAssessmentToCourse(user, TeacherAssessmentName, course)
	if err != nil {
		return
	}

	err = recipes.SearchUserAndShareCourse(user, course, lib.ApiUser2, lib.StudentAcl)
	if err != nil {
		return
	}

	err = publishAssessment(user, assessment)
	if err != nil {
		return
	}

	var Student response.User
	Student, err = recipes.SignIn(lib.ApiUser2, lib.Password)
	if err != nil {
		return
	}

	log.Printf("Assign grade to %s assessment for %s", assessment.Name, lib.ApiUser2)
	recipes.SleepBefore()
	err = assignAssessmentGrade(user, assessment, Student)
	if err != nil {
		return
	}
	log.Printf(".Grade assigned successfully to %s", lib.ApiUser2)
	recipes.SleepAfter()

	log.Printf("Publish grade for %s", Student.Email)
	recipes.SleepBefore()
	err = publishAssessmentGrade(user, assessment, Student)
	if err != nil {
		return
	}
	log.Printf(".Grade published successfully for %s", lib.ApiUser2)
	recipes.SleepAfter()
}

func publishAssessment(user response.User, assessment response.Assessment) error {
	_, err := assessments.UpdateAssessmentCompletionStatus(
		user.JwtToken,
		request.UpdateAssessmentStatusReqBody{Completed: true},
		common.ResourceIdParam{
			AssessmentId: assessment.ID,
			CourseId:     assessment.Course.ID,
			KeyId:        assessment.Key.ID,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func assignAssessmentGrade(user response.User, assessment response.Assessment, student response.User) error {
	resGradingSystems, _ := gradingSystems.GetGradingSystems(user.JwtToken, gradingSystems.GetGradingSystemsParam{
		IncludeCounts: false,
		ExcludeEmpty:  false,
	})

	var alphabeticGradingSystem response.GradingSystem
	for _, gradingSystem := range resGradingSystems {
		if *gradingSystem.Type == lib.AlphabeticGradingSystemType {
			alphabeticGradingSystem = gradingSystem
			break
		}
	}

	assessmentId := assessment.ID
	courseId := assessment.Course.ID
	err := assessments.AddGradingSystemToAssessment(user.JwtToken, request.GradingSystemIdParam{
		GradingSystemId: alphabeticGradingSystem.ID,
		AssessmentId:    &assessmentId,
		CourseId:        &courseId,
		KeyId:           assessment.Key.ID,
	})
	if err != nil {
		return err
	}

	_, err = teacherKeys.AssignAssessmentGradeForAStudentAsTeacher(
		user.JwtToken,
		request.UpdateGradeReqBody{Grade: alphabeticGradingSystem.Grades[0]},
		request.ClassroomIdParam{
			StudentId: student.ID,
			ResourceIds: common.ResourceIdParam{
				AssessmentId: assessmentId,
				CourseId:     courseId,
				KeyId:        assessment.Key.ID,
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func publishAssessmentGrade(user response.User, assessment response.Assessment, student response.User) error {
	err := teacherKeys.BulkPublishAssessmentGradesForAStudent(
		user.JwtToken,
		teacherKeys.PublishStudentGradeReqBody{AssessmentIds: assessment.ID},
		request.ClassroomIdParam{
			StudentId: student.ID,
			ResourceIds: common.ResourceIdParam{
				CourseId: assessment.Course.ID,
				KeyId:    assessment.Key.ID,
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
