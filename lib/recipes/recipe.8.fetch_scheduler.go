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
	SchedulerKeyName   = "Tuition Class"
	SchedulerBlockName = "Final Exam"
	SchedulerPodName   = "Math Quiz"
	DueDate            = "2023-03-10T14:19:04.027Z"
)

func FetchScheduler() {
	log.Info("Objective: Set Due Date for Course and Pod, and Fetch Scheduler Events")
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
	log.Info("Set due date for block")
	block, err := setBlockDueDate(user, key)
	if err != nil {
		return
	}
	log.Info(".Course due date set successfully")

	log.Info("Set due date for pod")
	err = setPodDueDate(user, block)
	if err != nil {
		return
	}
	log.Info(".Pod due date set successfully")

	log.Info("Show due date events")
	err = fetchSchedulerEvents(user)
	if err != nil {
		return
	}
	log.Info(".Displayed block & pod due date events")
}

func setBlockDueDate(user response.User, key response.Key) (response.Course, error) {
	block, err := recipes.AddCourse(user, SchedulerBlockName, key)
	if err != nil {
		return block, err
	}
	dueDate := DueDate
	block, err = courses.UpdateCourse(
		user.JwtToken,
		courses.UpdateCourseReqBody{DueDate: &dueDate},
		common.ResourceIdParam{BlockId: block.ID, KeyId: block.Key.ID})
	if err != nil {
		return block, err
	}
	return block, nil
}

func setPodDueDate(user response.User, block response.Course) error {
	pod, err := recipes.AddPodToBlock(user, SchedulerPodName, block)
	if err != nil {
		return err
	}
	dueDate := DueDate
	_, err = assessments.UpdateAssessment(
		user.JwtToken,
		request.UpdateAssessmentReqBody{DueDate: &dueDate},
		common.ResourceIdParam{PodId: pod.ID, KeyId: pod.Key.ID})
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
	for _, blockEvent := range allEvents.DueDateEvent.Courses {
		log.Printf(".Course %s is due on %s", blockEvent.Name, *blockEvent.DueDate)
	}
	for _, podEvent := range allEvents.DueDateEvent.Assessments {
		log.Printf(".Pod %s is due on %s", podEvent.Name, podEvent.DueDate)
	}
	return nil
}
