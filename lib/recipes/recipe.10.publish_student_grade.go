package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/assessments/assessments.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/scales"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	teacherKeys "github.com/snowpal/pitch-classroom-sdk/lib/endpoints/teacher_keys/teacher_keys.2.teachers"

	log "github.com/sirupsen/logrus"
)

const (
	TeacherKeyName   = "School"
	TeacherBlockName = "Math"
	TeacherPodName   = "Final Exam"
)

func PublishStudentGrade() {
	log.Info("Objective: Assign and Publish Assessment Grades for a Student")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	var user response.User
	user, err = recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	var key response.Key
	key, err = recipes.AddTeacherKey(user, TeacherKeyName)
	if err != nil {
		return
	}

	var block response.Course
	block, err = recipes.AddCourse(user, TeacherBlockName, key)
	if err != nil {
		return
	}

	var pod response.Assessment
	pod, err = recipes.AddPodToBlock(user, TeacherPodName, block)
	if err != nil {
		return
	}

	err = recipes.SearchUserAndShareBlock(user, block, "api_read_user", lib.ReadAcl)
	if err != nil {
		return
	}

	err = publishAssessment(user, pod)
	if err != nil {
		return
	}

	var readUser response.User
	readUser, err = recipes.SignIn(lib.ReadUser, lib.Password)
	if err != nil {
		return
	}

	log.Printf("Assign grade to %s pod for %s", pod.Name, lib.ReadUser)
	recipes.SleepBefore()
	err = assignAssessmentGrade(user, pod, readUser)
	if err != nil {
		return
	}
	log.Printf(".Grade assigned successfully to %s", lib.ReadUser)
	recipes.SleepAfter()

	log.Printf("Publish grade for %s", readUser.Email)
	recipes.SleepBefore()
	err = publishAssessmentGrade(user, pod, readUser)
	if err != nil {
		return
	}
	log.Printf(".Grade published successfully for %s", lib.ReadUser)
	recipes.SleepAfter()
}

func publishAssessment(user response.User, pod response.Assessment) error {
	_, err := assessments.UpdateAssessmentCompletionStatus(
		user.JwtToken,
		request.UpdateAssessmentStatusReqBody{Completed: true},
		common.ResourceIdParam{
			AssessmentId: pod.ID,
			CourseId:     pod.Course.ID,
			KeyId:        pod.Key.ID,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func assignAssessmentGrade(user response.User, pod response.Assessment, student response.User) error {
	resScales, _ := scales.GetScales(user.JwtToken, scales.GetScalesParam{
		IncludeCounts: false,
		ExcludeEmpty:  true,
	})
	var alphabeticScale response.Scale
	for _, scale := range resScales {
		if *scale.Type == lib.AlphabeticScaleType {
			alphabeticScale = scale
			break
		}
	}
	podId := pod.ID
	blockId := pod.Course.ID
	err := assessments.AddScaleToAssessment(user.JwtToken, request.ScaleIdParam{
		ScaleId:      alphabeticScale.ID,
		AssessmentId: &podId,
		CourseId:     &blockId,
		KeyId:        pod.Key.ID,
	})
	if err != nil {
		return err
	}
	_, err = teacherKeys.AssignAssessmentGradeForAStudentAsTeacher(
		user.JwtToken,
		request.UpdateScaleValueReqBody{ScaleValue: alphabeticScale.ScaleValues[0]},
		request.ClassroomIdParam{
			StudentId: student.ID,
			ResourceIds: common.ResourceIdParam{
				AssessmentId: podId,
				CourseId:     blockId,
				KeyId:        pod.Key.ID,
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func publishAssessmentGrade(user response.User, pod response.Assessment, student response.User) error {
	err := teacherKeys.BulkPublishAssessmentGradesForAStudent(
		user.JwtToken,
		teacherKeys.PublishStudentGradeReqBody{PodIds: pod.ID},
		request.ClassroomIdParam{
			StudentId: student.ID,
			ResourceIds: common.ResourceIdParam{
				CourseId: pod.Course.ID,
				KeyId:    pod.Key.ID,
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
