package response

type ProblemDetails struct {
	Type   string `json:"type,omitempty"`
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func NewProblemDetails(problemDetailsType string, status int, title string, detail string) ProblemDetails {
	return ProblemDetails{
		Type:   problemDetailsType,
		Status: status,
		Title:  title,
		Detail: detail,
	}
}
