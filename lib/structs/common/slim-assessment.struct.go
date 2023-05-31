package common

type SlimAssessment struct {
	ID      string       `json:"id"`
	Name    string       `json:"assessmentName"`
	Key     *SlimKey     `json:"key"`
	Course  *SlimCourse  `json:"course"`
	Courses []SlimCourse `json:"courses"`
}
