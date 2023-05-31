package request

type NoteReqBody struct {
	NoteText string `json:"noteText"`
}

type NoteIdParam struct {
	KeyId        string
	CourseId     *string
	AssessmentId *string
	NoteId       *string
}
