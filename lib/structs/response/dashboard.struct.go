package response

import (
	common2 "github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

type Dashboard struct {
	Dashboard DashboardResources `json:"dashboard"`
}

type DashboardResources struct {
	RecentlyModifiedResources *RecentlyModifiedResources `json:"recentlyModified"`
	ShortlyDueResources       *DueShortlyResources       `json:"dueShortly"`
}

type RecentlyModifiedKeys struct {
	Keys []common2.SlimKey `json:"keys"`
}

type RecentlyModifiedResources struct {
	Courses     []DashboardCourse     `json:"courses"`
	Assessments []DashboardAssessment `json:"assessments"`
}

type DueShortlyResources struct {
	Courses     *[]DashboardCourse    `json:"courses"`
	Assessments []DashboardAssessment `json:"assessments"`
	Tasks       []DashboardTask       `json:"tasks"`
}

type DashboardCourse struct {
	ID      string `json:"id"`
	Name    string `json:"courseName"`
	DueDate string `json:"courseDueDate"`

	Key        *common2.SlimKey `json:"key"`
	CourseType *CourseType      `json:"courseType"`

	Creator  common2.ResourceCreator  `json:"creator"`
	Modifier common2.ResourceModifier `json:"modifier"`
}

type DashboardAssessment struct {
	ID      string `json:"id"`
	Name    string `json:"assessmentName"`
	DueDate string `json:"assessmentDueDate"`

	Key     *common2.SlimKey            `json:"key"`
	Courses *[]CourseWithAssessmentType `json:"courses"`

	Creator  common2.ResourceCreator  `json:"creator"`
	Modifier common2.ResourceModifier `json:"modifier"`
}

type CourseWithAssessmentType struct {
	ID             string           `json:"id"`
	Name           string           `json:"courseName"`
	Key            *common2.SlimKey `json:"key"`
	AssessmentType *AssessmentType  `json:"assessmentType"`
}

type DashboardTask struct {
	ID      string `json:"id"`
	Name    string `json:"taskName"`
	DueDate string `json:"taskDueDate"`

	Key        *common2.SlimKey        `json:"key"`
	Course     *common2.SlimCourse     `json:"course"`
	Assessment *common2.SlimAssessment `json:"assessment"`
	Courses    *[]common2.SlimCourse   `json:"courses"`

	Creator  common2.ResourceCreator  `json:"creator"`
	Modifier common2.ResourceModifier `json:"modifier"`
}

type DashboardUnreadCount struct {
	DueTasks       int `json:"dueTasks"`
	DueCourses     int `json:"dueCourses"`
	DueAssessments int `json:"dueAssessments"`

	Notifications int `json:"notifications"`
	Conversations int `json:"conversations"`
}
