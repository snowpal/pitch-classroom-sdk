package response

import (
	common2 "github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

type SearchResources struct {
	Results []SearchResource `json:"results"`
}

type SearchResource struct {
	ID   string `json:"id"`
	Type string `json:"type"`

	// Relation attribute
	IsRelated bool `json:"isRelated"`

	KeyName    *string `json:"keyName"`
	KeyType    *string `json:"keyType"`
	Coursename *string `json:"courseName"`
	PodName    *string `json:"assessmentName"`

	Key     *common2.SlimKey      `json:"key"`
	Course  *common2.SlimCourse   `json:"course"`
	Courses *[]common2.SlimCourse `json:"courses"`

	Modifier common2.ResourceModifier `json:"modifier"`
}

type Relations struct {
	Relationships Relationships `json:"relationships"`
}

type Relationships struct {
	Keys        []common2.SlimKey        `json:"keys"`
	Courses     []common2.SlimCourse     `json:"courses"`
	Assessments []common2.SlimAssessment `json:"pods"`
}
