package response

import (
	common2 "github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

type AssessmentTypes struct {
	AssessmentTypes []AssessmentType `json:"assessmentTypes"`
}

type AssessmentType struct {
	ID          string                    `json:"id"`
	Name        string                    `json:"assessmentTypeName"`
	Assessments *[]common2.SlimAssessment `json:"assessments"`

	TeacherReadOnly *bool `json:"teacherReadOnly"`

	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}
