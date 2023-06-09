package response

import (
	common2 "github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

type Courses struct {
	Courses []Course `json:"courses"`
}

type Course struct {
	ID                string  `json:"id"`
	Name              string  `json:"courseName"`
	CourseId          string  `json:"courseId"`
	Description       string  `json:"courseDescription"`
	SimpleDescription string  `json:"simpleDescription"`
	Color             string  `json:"color"`
	Tags              string  `json:"tags"`
	Grade             *string `json:"grade"`

	Attributes    []common2.DisplayAttribute `json:"attributes"`
	CourseType    *CourseType                `json:"courseType"`
	GradingSystem *GradingSystem             `json:"gradingSystem"`
	TaggedUsers   []TaggedUser               `json:"taggedUsers"`
	Key           *common2.SlimKey           `json:"key"`

	// Boolean Attributes
	Completed  *bool `json:"completed"`
	Archived   *bool `json:"archived"`
	KanbanMode *bool `json:"kanbanMode"`
	CanUnlink  *bool `json:"canUnlink"`
	PublicKey  *bool `json:"publicKey"`

	// Time Attributes
	DueDate   string `json:"courseDueDate"`
	StartTime string `json:"courseStartTime"`
	EndTime   string `json:"courseEndTime"`

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
	AssessmentsCount *int `json:"assessmentsCount"`
	TasksCount       *int `json:"tasksCount"`
	ChecklistsCount  *int `json:"checklistsCount"`
	AttachmentsCount *int `json:"attachmentsCount"`
	CommentsCount    *int `json:"commentsCount"`
	NotesCount       *int `json:"notesCount"`

	Creator      common2.ResourceCreator  `json:"creator"`
	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}

type UpdateCourseGrade struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Grade                string `json:"grade"`
	NumericGradingSystem int    `json:"numericGradingSystem"`

	Key common2.SlimKey `json:"key"`
}
