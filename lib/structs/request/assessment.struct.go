package request

type AddAssessmentReqBody struct {
	Name string `json:"assessmentName"`
}

type UpdateAssessmentDescReqBody struct {
	Description   string  `json:"assessmentDescription"`
	TaggedUserIds *string `json:"taggedUserIds"`
}

type BulkArchiveAssessmentsReqBody struct {
	AssessmentIds string `json:"assessmentIds"`
}

type UpdateAssessmentStatusReqBody struct {
	Completed bool `json:"assessmentCompleted"`
}

type UpdateAssessmentReqBody struct {
	Name              *string `json:"assessmentName"`
	SimpleDescription *string `json:"simpleDescription"`
	DueDate           *string `json:"assessmentDueDate"`
	Color             *string `json:"assessmentColor"`
	Tags              *string `json:"assessmentTags"`
	KanbanMode        *bool   `json:"kanbanMode"`
}

type GetAssessmentsParam struct {
	KeyId      string
	CourseId   *string
	BatchIndex int
}

type AddAssessmentTypeIdParam struct {
	AssessmentId     string
	AssessmentTypeId string
	KeyId            string
	CourseId         *string
}

type AssessmentByTemplateParam struct {
	KeyId        string
	CourseId     *string
	TemplateId   string
	ExcludeTasks bool
}

type CopyMoveAssessmentParam struct {
	AssessmentId string
	KeyId        string
	TargetKeyId  string

	CourseId       string
	TargetCourseId string

	AllTasks      bool
	AllChecklists bool
}
