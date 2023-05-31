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
	Blocks       []UserBlock              `json:"courses"`
	Pods         []common2.SlimAssessment `json:"pods"`
	LastModified string                   `json:"lastModified"`
}

type UserBlock struct {
	ID           string                   `json:"id"`
	Name         string                   `json:"keyName"`
	Pods         []common2.SlimAssessment `json:"pods"`
	LastModified string                   `json:"lastModified"`
}

type FilteredKeys struct {
	Keys []FilteredKey `json:"keys"`
}

type BlocksAndPods struct {
	Blocks []common2.SlimCourse     `json:"courses"`
	Pods   []common2.SlimAssessment `json:"pods"`
}

type FilteredKey struct {
	ID               string        `json:"id"`
	Name             string        `json:"keyName"`
	Type             string        `json:"keyType"`
	CreatedByMe      BlocksAndPods `json:"createdByMe"`
	SharedWithMe     BlocksAndPods `json:"sharedWithMe"`
	SharedWithOthers BlocksAndPods `json:"sharedWithOthers"`
}

type BlockTypesKeys struct {
	Keys []BlockTypesKey `json:"keys"`
}

type BlockTypesKey struct {
	ID           string       `json:"id"`
	Name         string       `json:"keyName"`
	Type         string       `json:"keyType"`
	BlockTypes   []CourseType `json:"blockTypes"`
	LastModified string       `json:"lastModified"`
}

type PodTypesKeysKeyPod struct {
	KeyPod PodTypesKeys `json:"keyPod"`
}

type PodTypesKeysBlockPod struct {
	BlockPod PodTypesKeys `json:"blockPod"`
}

type PodTypesKeysOtherPod struct {
	OtherPod PodTypesKeys `json:"otherPod"`
}

type PodTypesKeys struct {
	Keys *[]PodTypesKey `json:"keys"`
	Key  *PodTypesKey   `json:"key"`
}

type PodTypesKey struct {
	ID              string            `json:"id"`
	Name            string            `json:"keyName"`
	Type            string            `json:"keyType"`
	AssessmentTypes *[]AssessmentType `json:"podTypes"`
	LastModified    string            `json:"lastModified"`
}

type ScalesKeysBlock struct {
	Course ScalesKeys `json:"course"`
}

type ScalesKeysKeyPod struct {
	KeyPod ScalesKeys `json:"keyPod"`
}

type ScalesKeysBlockPod struct {
	BlockPod ScalesKeys `json:"blockPod"`
}

type ScalesKeysOtherPod struct {
	OtherPod ScalesKeys `json:"otherPod"`
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

type TasksStatusBlock struct {
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
	Blocks     []UserBlock        `json:"courses"`
}

type BlockScaleValue struct {
	ID           string `json:"id"`
	Name         string `json:"courseName"`
	ScaleValue   string `json:"scaleValue"`
	NumericScale int    `json:"numericScale"`
}

type PodScaleValue struct {
	ID           string `json:"id"`
	Name         string `json:"assessmentName"`
	ScaleValue   string `json:"scaleValue"`
	NumericScale int    `json:"numericScale"`
}

type ScaleValues struct {
	Scale  Scale               `json:"scale"`
	Key    common2.SlimKey     `json:"key"`
	Course *common2.SlimCourse `json:"course"`
	Blocks *[]BlockScaleValue  `json:"courses"`
	Pods   []PodScaleValue     `json:"pods"`
}

type BlockGrade struct {
	ID       string          `json:"id"`
	Name     string          `json:"courseName"`
	Key      common2.SlimKey `json:"key"`
	Students []Student       `json:"students"`
}

type PodGrade struct {
	ID       string             `json:"id"`
	Name     string             `json:"assessmentName"`
	Key      common2.SlimKey    `json:"key"`
	Course   common2.SlimCourse `json:"course"`
	Students []Student          `json:"students"`
}
