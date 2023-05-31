package response

import (
	common2 "github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

type UserKeys struct {
	Keys []UserKey `json:"keys"`
}

type UserKey struct {
	ID           string                   `json:"id"`
	Name         string                   `json:"keyName"`
	Type         string                   `json:"keyType"`
	Courses      []UserCourse             `json:"courses"`
	Assessments  []common2.SlimAssessment `json:"assessments"`
	LastModified string                   `json:"lastModified"`
}

type UserCourse struct {
	ID           string                   `json:"id"`
	Name         string                   `json:"keyName"`
	Assessments  []common2.SlimAssessment `json:"assessments"`
	LastModified string                   `json:"lastModified"`
}

type FilteredKeys struct {
	Keys []FilteredKey `json:"keys"`
}

type CoursesAndAssessments struct {
	Courses     []common2.SlimCourse     `json:"courses"`
	Assessments []common2.SlimAssessment `json:"assessments"`
}

type FilteredKey struct {
	ID               string                `json:"id"`
	Name             string                `json:"keyName"`
	Type             string                `json:"keyType"`
	CreatedByMe      CoursesAndAssessments `json:"createdByMe"`
	SharedWithMe     CoursesAndAssessments `json:"sharedWithMe"`
	SharedWithOthers CoursesAndAssessments `json:"sharedWithOthers"`
}

type CourseTypesKeys struct {
	Keys []CourseTypesKey `json:"keys"`
}

type CourseTypesKey struct {
	ID           string       `json:"id"`
	Name         string       `json:"keyName"`
	Type         string       `json:"keyType"`
	CourseTypes  []CourseType `json:"courseTypes"`
	LastModified string       `json:"lastModified"`
}

type AssessmentTypesKeysAssessment struct {
	Assessment AssessmentTypesKeys `json:"assessment"`
}

type AssessmentTypesKeys struct {
	Keys *[]AssessmentTypesKey `json:"keys"`
	Key  *AssessmentTypesKey   `json:"key"`
}

type AssessmentTypesKey struct {
	ID              string            `json:"id"`
	Name            string            `json:"keyName"`
	Type            string            `json:"keyType"`
	AssessmentTypes *[]AssessmentType `json:"assessmentTypes"`
	LastModified    string            `json:"lastModified"`
}

type ScalesKeysCourse struct {
	Course ScalesKeys `json:"course"`
}

type ScalesKeysAssessment struct {
	Assessment ScalesKeys `json:"assessment"`
}

type ScalesKeys struct {
	Keys *[]ScalesKey `json:"keys"`
	Key  *ScalesKey   `json:"key"`
}

type ScalesKey struct {
	ID           string   `json:"id"`
	Name         string   `json:"keyName"`
	Type         string   `json:"keyType"`
	Scales       *[]Scale `json:"scales"`
	LastModified string   `json:"lastModified"`
}

type TaskStatus struct {
	Complete   int `json:"complete"`
	Incomplete int `json:"incomplete"`
}

type TasksStatusKeys struct {
	Keys []TasksStatusKey `json:"keys"`
}

type TasksStatusKey struct {
	ID           string     `json:"id"`
	Name         string     `json:"keyName"`
	TaskStatus   TaskStatus `json:"taskStatus"`
	LastModified string     `json:"lastModified"`
}

type TasksStatusCourse struct {
	ID           string          `json:"id"`
	Name         string          `json:"courseName"`
	TaskStatus   TaskStatus      `json:"taskStatus"`
	Key          common2.SlimKey `json:"key"`
	LastModified string          `json:"lastModified"`
}

type LinkedResourcesKey struct {
	SlimKey common2.SlimKey `json:"key"`
}

type LinkedResources struct {
	CurrentKey LinkedResourcesKey `json:"currentKey"`
	SharedKey  LinkedResourcesKey `json:"sharedKey"`
	Keys       *[]UserKey         `json:"keys"`
	Courses    []UserCourse       `json:"courses"`
}

type CourseScaleValue struct {
	ID           string `json:"id"`
	Name         string `json:"courseName"`
	ScaleValue   string `json:"scaleValue"`
	NumericScale int    `json:"numericScale"`
}

type AssessmentScaleValue struct {
	ID           string `json:"id"`
	Name         string `json:"assessmentName"`
	ScaleValue   string `json:"scaleValue"`
	NumericScale int    `json:"numericScale"`
}

type ScaleValues struct {
	Scale       Scale                  `json:"scale"`
	Key         common2.SlimKey        `json:"key"`
	Course      *common2.SlimCourse    `json:"course"`
	Courses     *[]CourseScaleValue    `json:"courses"`
	Assessments []AssessmentScaleValue `json:"assessments"`
}

type CourseGrade struct {
	ID       string          `json:"id"`
	Name     string          `json:"courseName"`
	Key      common2.SlimKey `json:"key"`
	Students []Student       `json:"students"`
}

type AssessmentGrade struct {
	ID       string             `json:"id"`
	Name     string             `json:"assessmentName"`
	Key      common2.SlimKey    `json:"key"`
	Course   common2.SlimCourse `json:"course"`
	Students []Student          `json:"students"`
}
