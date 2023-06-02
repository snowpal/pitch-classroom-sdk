package request

type KeyToKeyRelationParam struct {
	KeyId       string
	TargetKeyId string
}

type KeyToCourseRelationParam struct {
	KeyId          string
	TargetCourseId string
}

type KeyToAssessmentRelationParam struct {
	KeyId string

	TargetAssessmentId string
	TargetKeyId        string
	TargetCourseId     string
}

type CourseToCourseRelationParam struct {
	CourseId       string
	TargetCourseId string
}

type CourseToAssessmentRelationParam struct {
	CourseId string

	TargetAssessmentId string
	TargetKeyId        string
	TargetCourseId     string
}

type AssessmentToAssessmentRelationParam struct {
	AssessmentId   string
	SourceKeyId    string
	SourceCourseId string

	TargetAssessmentId string
	TargetKeyId        string
	TargetCourseId     string
}
