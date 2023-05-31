package common

type SlimCourse struct {
	ID   string   `json:"id"`
	Name string   `json:"courseName"`
	Key  *SlimKey `json:"key"`
}
