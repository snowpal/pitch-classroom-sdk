package recipes

import (
	"github.com/snowpal/pitch-classroom-sdk/lib"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/assessments/assessments.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/courses/courses.1"
	"github.com/snowpal/pitch-classroom-sdk/lib/endpoints/scheduler"
	"github.com/snowpal/pitch-classroom-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/request"
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

const (
	SchedulerKeyName        = "Tuition Class"
	SchedulerCourseName     = "Final Exam"
	SchedulerAssessmentName = "Math Quiz"
	DueDate                 = "2023-03-10T14:19:04.027Z"
)

func FetchScheduler() {
	log.Info("Objective: Set Due Date for Course and Assessment, and Fetch Scheduler Events")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	var key response.Key
	key, _ = recipes.AddTeacherKey(user, SchedulerKeyName)
	log.Info("Set due date for course")
	course, err := setCourseDueDate(user, key)
	if err != nil {
		return
	}
	log.Info(".Course due date set successfully")

	log.Info("Set due date for assessment")
	err = setAssessmentDueDate(user, course)
	if err != nil {
		return
	}
	log.Info(".Assessment due date set successfully")

	log.Info("Show due date events")
	err = fetchSchedulerEvents(user)
	if err != nil {
		return
	}
	log.Info(".Displayed course & assessment due date events")
}

func setCourseDueDate(user response.User, key response.Key) (response.Course, error) {
	course, err := recipes.AddCourse(user, SchedulerCourseName, key)
	if err != nil {
		return course, err
	}
	dueDate := DueDate
	course, err = courses.UpdateCourse(
		user.JwtToken,
		courses.UpdateCourseReqBody{DueDate: &dueDate},
		common.ResourceIdParam{CourseId: course.ID, KeyId: course.Key.ID})
	if err != nil {
		return course, err
	}
	return course, nil
}

func setAssessmentDueDate(user response.User, course response.Course) error {
	assessment, err := recipes.AddAssessmentToCourse(user, SchedulerAssessmentName, course)
	if err != nil {
		return err
	}
	dueDate := DueDate
	_, err = assessments.UpdateAssessment(
		user.JwtToken,
		request.UpdateAssessmentReqBody{DueDate: &dueDate},
		common.ResourceIdParam{AssessmentId: assessment.ID, KeyId: assessment.Key.ID})
	if err != nil {
		return err
	}
	return nil
}

func fetchSchedulerEvents(user response.User) error {
	allEvents, err := scheduler.GetEventsForGivenDay(user.JwtToken, request.EventDateParam{StartDate: DueDate})
	if err != nil {
		return err
	}
	for _, courseEvent := range allEvents.DueDateEvent.Courses {
		log.Printf(".Course %s is due on %s", courseEvent.Name, *courseEvent.DueDate)
	}
	for _, assessmentEvent := range allEvents.DueDateEvent.Assessments {
		log.Printf(".Assessment %s is due on %s", assessmentEvent.Name, assessmentEvent.DueDate)
	}
	return nil
}
