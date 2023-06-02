package response

import (
	common2 "github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

type CourseTypes struct {
	CourseTypes []CourseType `json:"courseTypes"`
}

type CourseType struct {
	ID      string                `json:"id"`
	Name    string                `json:"courseTypeName"`
	Courses *[]common2.SlimCourse `json:"courses"`

	TeacherReadOnly *bool `json:"teacherReadOnly"`

	Modifier     *common2.ResourceModifier `json:"modifier"`
	LastModified string                    `json:"lastModified"`
}
