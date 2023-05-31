package request

type CommentReqBody struct {
	CommentText   string  `json:"commentText"`
	TaggedUserIds *string `json:"taggedUserIds"`
}

type CommentIdParam struct {
	KeyId        string
	CourseId     *string
	AssessmentId *string
	CommentId    *string
}
