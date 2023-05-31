package request

type KeyToKeyRelationParam struct {
	KeyId       string
	TargetKeyId string
}

type KeyToCourseRelationParam struct {
	KeyId         string
	TargetBlockId string
}

type KeyToAssessmentRelationParam struct {
	KeyId string

	TargetPodId   string
	TargetKeyId   string
	TargetBlockId string
}

type CourseToCourseRelationParam struct {
	BlockId       string
	TargetBlockId string
}

type CourseToAssessmentRelationParam struct {
	BlockId string

	TargetPodId   string
	TargetKeyId   string
	TargetBlockId string
}

type AssessmentToAssessmentRelationParam struct {
	PodId         string
	SourceKeyId   string
	SourceBlockId string

	TargetPodId   string
	TargetKeyId   string
	TargetBlockId string
}
