package response

import (
	common2 "github.com/snowpal/pitch-classroom-sdk/lib/structs/common"
)

type Comments struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID          string       `json:"id"`
	CommentText string       `json:"commentText"`
	TaggedUsers []TaggedUser `json:"taggedUsers"`

	CanEdit   *bool `json:"canEdit"`
	CanDelete *bool `json:"canDelete"`

	Key        *common2.SlimKey        `json:"key"`
	Course     *common2.SlimCourse     `json:"course"`
	Assessment *common2.SlimAssessment `json:"assessment"`

	Creator      common2.ResourceCreator  `json:"creator"`
	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}

type RecentComments struct {
	Comments []RecentComment `json:"comments"`
}

type RecentComment struct {
	ID          string `json:"id"`
	CommentText string `json:"commentText"`

	Key        *common2.SlimKey        `json:"key"`
	Course     *common2.SlimCourse     `json:"course"`
	Courses    *[]common2.SlimCourse   `json:"courses"`
	Assessment *common2.SlimAssessment `json:"assessment"`

	StudentId *string `json:"studentId"`

	Creator      common2.ResourceCreator  `json:"creator"`
	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}
