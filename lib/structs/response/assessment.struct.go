package response

import (
	common2 "github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

type Assessments struct {
	Assessments []Assessment `json:"assessments"`
}

type Assessment struct {
	ID                string  `json:"id"`
	Name              string  `json:"assessmentName"`
	Description       string  `json:"assessmentDescription"`
	SimpleDescription string  `json:"simpleDescription"`
	Color             string  `json:"color"`
	Tags              string  `json:"tags"`
	ScaleValue        *string `json:"scaleValue"`

	Attributes     []common2.DisplayAttribute `json:"attributes"`
	AssessmentType *AssessmentType            `json:"assessmentType"`
	Scale          *Scale                     `json:"scale"`
	TaggedUsers    []TaggedUser               `json:"taggedUsers"`
	Key            *common2.SlimKey           `json:"key"`
	Course         *common2.SlimCourse        `json:"course"`

	// Boolean Attributes
	Completed    *bool `json:"completed"`
	Archived     *bool `json:"archived"`
	KanbanMode   *bool `json:"kanbanMode"`
	CanUnlink    *bool `json:"canUnlink"`
	CanUnpublish bool  `json:"canUnpublish"`
	PublicKey    *bool `json:"publicKey"`

	// Time Attributes
	DueDate string `json:"assessmentDueDate"`

	// Acl Attributes
	Acl            *string       `json:"acl"`
	IsShared       *bool         `json:"isShared"`
	CanLeave       *bool         `json:"canLeave"`
	AllowDelete    *bool         `json:"allowDelete"`
	CanDelete      *bool         `json:"canDelete"`
	SharedUsers    *[]SharedUser `json:"sharedUsers"`
	CurrentUserAcl SharedUser    `json:"currentUserAcl"`

	// Count Attributes
	KeysCount        *int `json:"keysCount"`
	CoursesCount     *int `json:"coursesCount"`
	TasksCount       *int `json:"tasksCount"`
	ChecklistsCount  *int `json:"checklistsCount"`
	AttachmentsCount *int `json:"attachmentsCount"`
	CommentsCount    *int `json:"commentsCount"`
	NotesCount       *int `json:"notesCount"`

	Creator      common2.ResourceCreator  `json:"creator"`
	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}

type UpdateAssessmentScaleValue struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ScaleValue   string `json:"scaleValue"`
	NumericScale int    `json:"numericScale"`

	Key    common2.SlimKey     `json:"key"`
	Course *common2.SlimCourse `json:"course"`
}
