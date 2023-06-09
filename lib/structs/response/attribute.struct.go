package response

type ResourceAttribute struct {
	Name    string `json:"attributeName"`
	CanHide bool   `json:"canHide"`
}

type ResourceAttributes struct {
	KeyAttributes        []ResourceAttribute
	CourseAttributes     []ResourceAttribute
	AssessmentAttributes []ResourceAttribute
}
