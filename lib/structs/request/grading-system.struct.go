package request

type GradingSystemReqBody struct {
	Name   string   `json:"gradingSystemName"`
	Type   string   `json:"gradingSystemType"`
	Grades []string `json:"grades"`
}

type UpdateGradeReqBody struct {
	Grade string `json:"grade"`
}

type GradingSystemIdParam struct {
	GradingSystemId string
	KeyId           string
	CourseId        *string
	AssessmentId    *string
}
