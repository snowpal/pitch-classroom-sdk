package request

type AddAssessmentReqBody struct {
	Name string `json:"assessmentName"`
}

type UpdateAssessmentDescReqBody struct {
	Description   string  `json:"assessmentDescription"`
	TaggedUserIds *string `json:"taggedUserIds"`
}

type BulkArchiveAssessmentsReqBody struct {
	PodIds string `json:"assessmentIds"`
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

type AssessmentAclReqBody struct {
	Acl string `json:"assessmentAcl"`
}

type AssessmentBulkShareReqBody struct {
	Acl    string `json:"assessmentAcl"`
	PodIds string `json:"assessmentIds"`
}

type GetAssessmentsParam struct {
	KeyId      string
	BlockId    *string
	BatchIndex int
}

type AddPodTypeIdParam struct {
	PodId     string
	PodTypeId string
	KeyId     string
	BlockId   *string
}

type PodByTemplateParam struct {
	KeyId        string
	BlockId      *string
	TemplateId   string
	ExcludeTasks bool
}

type CopyMovePodParam struct {
	PodId       string
	KeyId       string
	TargetKeyId string

	BlockId       string
	TargetBlockId string

	AllTasks      bool
	AllChecklists bool
}
