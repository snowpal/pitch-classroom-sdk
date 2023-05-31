package request

type AddCourseReqBody struct {
	Name string `json:"courseName"`
}

type CourseAclReqBody struct {
	Acl string `json:"courseAcl"`
}

type GetCoursesParam struct {
	KeyId            string
	BatchIndex       int
	WriteOrHigherAcl bool
	Filter           string
}

type CopyMoveCourseParam struct {
	CourseId    string
	KeyId       string
	TargetKeyId string

	AssessmentIds  []string
	AllAssessments bool
	AllTasks       bool
	AllChecklists  bool
}
