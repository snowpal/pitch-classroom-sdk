package response

import (
	common2 "github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID        string                   `json:"id"`
	Name      string                   `json:"taskName"`
	DueDate   string                   `json:"taskDueDate"`
	Completed bool                     `json:"isCompleted"`
	Assignees []TaggedUser             `json:"assignees"`
	Key       *common2.SlimKey         `json:"key"`
	Course    *common2.SlimCourse      `json:"course"`
	Creator   common2.ResourceCreator  `json:"creator"`
	Modifier  common2.ResourceModifier `json:"modifier"`
}
