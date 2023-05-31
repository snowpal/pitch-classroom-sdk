package response

// Note(anish,3,03/06/2023) Most of the attributes are common to these templates, so the gateway should send a nested
// structure ideally.

type KeyTemplates struct {
	Templates []KeyTemplatesWithType `json:"templates"`
}

type KeyTemplatesWithType struct {
	Type      string        `json:"type"`
	Templates []KeyTemplate `json:"templates"`
}

type KeyTemplate struct {
	ID           string `json:"id"`
	Name         string `json:"templateName"`
	Description  string `json:"templateDescription"`
	Color        string `json:"color"`
	Tags         string `json:"tags"`
	PreviewUrl   string `json:"previewUrl"`
	Archived     bool   `json:"archived"`
	LastModified string `json:"lastModified"`
}

type CourseTemplates struct {
	Templates []CourseTemplate `json:"templates"`
}

type CourseTemplate struct {
	ID            string  `json:"id"`
	Name          string  `json:"templateName"`
	Description   string  `json:"templateDescription"`
	Color         string  `json:"color"`
	Tags          string  `json:"tags"`
	PreviewUrl    string  `json:"previewUrl"`
	Archived      bool    `json:"archived"`
	KeyTemplateId *string `json:"keyTemplateId"`
	LastModified  string  `json:"lastModified"`
}

type AssessmentTemplates struct {
	Templates []AssessmentTemplate `json:"templates"`
}

type AssessmentTemplate struct {
	ID               string  `json:"id"`
	Name             string  `json:"templateName"`
	Description      string  `json:"templateDescription"`
	Color            string  `json:"color"`
	Tags             string  `json:"tags"`
	PreviewUrl       string  `json:"previewUrl"`
	Archived         bool    `json:"archived"`
	KeyTemplateId    *string `json:"keyTemplateId"`
	CourseTemplateId *string `json:"courseTemplateId"`
	LastModified     string  `json:"lastModified"`
}
