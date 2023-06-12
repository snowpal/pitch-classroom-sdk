package response

import (
	"github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

type GradingSystems struct {
	GradingSystems []GradingSystem `json:"gradingSystems"`
}

type GradingSystem struct {
	ID              string                  `json:"id"`
	Name            string                  `json:"gradingSystemName"`
	Type            *string                 `json:"gradingSystemType"`
	Grades          []string                `json:"grades"`
	TeacherReadOnly *bool                   `json:"teacherReadOnly"`
	Modifier        common.ResourceModifier `json:"modifier"`
	LastModified    string                  `json:"lastModified"`
}
