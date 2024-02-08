package response

import "net/http"

type ProblemDetails struct {
	Type   string `json:"type,omitempty"`
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func NewProblemDetails() ProblemDetails {
	return ProblemDetails{}
}

func (p ProblemDetails) WithType(t string) ProblemDetails {
	p.Type = t

	return p
}

func (p ProblemDetails) WithStatus(s int) ProblemDetails {
	p.Status = s

	return p
}

func (p ProblemDetails) WithTitle(t string) ProblemDetails {
	p.Title = t

	return p
}

func (p ProblemDetails) WithDetail(d string) ProblemDetails {
	p.Detail = d

	return p
}

func (p ProblemDetails) WithInternalError() ProblemDetails {
	p.Type = "INTERNAL_ERROR"
	p.Status = http.StatusInternalServerError
	p.Title = "Internal error"
	p.Detail = "The server encountered an unexpected condition that prevented it from fulfilling the request."

	return p
}
