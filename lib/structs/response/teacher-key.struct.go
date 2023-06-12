package response

import (
	common2 "github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

type StudentGradeForCourseAndAssessment struct {
	ID           string                       `json:"id"`
	Name         string                       `json:"courseName"`
	Key          common2.SlimKey              `json:"key"`
	Assessment   *common2.SlimAssessment      `json:"assessment"`
	StudentGrade *StudentGrade                `json:"grade"`
	Assessments  *[]StudentGradeForAssessment `json:"assessments"`
	Students     *[]Student                   `json:"students"`
}

type StudentGradeForAssessment struct {
	ID           string        `json:"id"`
	Name         string        `json:"assessmentName"`
	StudentGrade *StudentGrade `json:"grade"`
}

type StudentGrade struct {
	Grade                string `json:"grade"`
	NumericGradingSystem int    `json:"numericGradingSystem"`
	Published            bool   `json:"published"`
	PublishedOn          string `json:"publishedOn"`
}

type Students struct {
	Students []Student `json:"students"`
}

type Student struct {
	ID            string              `json:"id"`
	ProfileID     string              `json:"profileId"`
	Email         string              `json:"email"`
	Username      string              `json:"username"`
	FirstName     string              `json:"firstName"`
	MiddleName    string              `json:"middleName"`
	LastName      string              `json:"lastName"`
	PhoneNumber   string              `json:"phoneNumber"`
	AddressUserBy string              `json:"addressUserBy"`
	UserInitial   string              `json:"userInitial"`
	AvatarName    string              `json:"avatarName"`
	AvatarUrl     string              `json:"avatarUrl"`
	Coursename    string              `json:"courseName"`
	StudentGrade  *StudentGrade       `json:"grade"`
	Key           *common2.SlimKey    `json:"key"`
	Course        *common2.SlimCourse `json:"course"`
}
