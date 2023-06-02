package response

import (
	common2 "github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

type AllEvents struct {
	DueDateEvent    *DueDateEvent    `json:"dueDate"`
	EndDateEvent    *EndDateEvent    `json:"endDate"`
	SchedulerEvents []SchedulerEvent `json:"schedulerEvents"`
}

type EndDateEvent struct {
	Courses []CourseEvent `json:"courses"`
}

type DueDateEvent struct {
	Tasks       TasksEvent        `json:"tasks"`
	Courses     []CourseEvent     `json:"courses"`
	Assessments []AssessmentEvent `json:"assessments"`
}

type CourseEvent struct {
	ID          string `json:"id"`
	Name        string `json:"courseName"`
	Description string `json:"courseDescription"`

	DueDate   *string `json:"courseDueDate"`
	StartTime *string `json:"courseStartTime"`
	EndTime   *string `json:"courseEndTime"`

	Key common2.SlimKey `json:"key"`
}

type AssessmentEvent struct {
	ID      string `json:"id"`
	Name    string `json:"assessmentName"`
	DueDate string `json:"assessmentDueDate"`

	Key    common2.SlimKey    `json:"key"`
	Course common2.SlimCourse `json:"course"`
}

type TasksEvent struct {
	KeyTasks        []TaskEvent `json:"keys"`
	CourseTasks     []TaskEvent `json:"courses"`
	AssessmentTasks []TaskEvent `json:"assessment"`
}

type TaskEvent struct {
	ID         string                  `json:"id"`
	Name       string                  `json:"taskName"`
	DueDate    string                  `json:"taskDueDate"`
	Key        common2.SlimKey         `json:"key"`
	Course     *common2.SlimCourse     `json:"course"`
	Assessment *common2.SlimAssessment `json:"assessment"`
}

type SchedulerEvent struct {
	SchedulerId      string            `json:"schedulerId"`
	StandaloneEvents []StandaloneEvent `json:"standaloneEvents"`
}

type StandaloneEvent struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	StartTime   string `json:"eventStartTime"`
	EndTime     string `json:"eventEndTime"`

	Creator  *common2.ResourceCreator  `json:"creator"`
	Modifier *common2.ResourceModifier `json:"modifier"`
}
