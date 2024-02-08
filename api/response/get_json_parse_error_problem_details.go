package response

import "net/http"

func GetJSONParseErrorProblemDetails() ProblemDetails {
	return NewProblemDetails().
		WithType("JSON_PARSE_ERROR").
		WithStatus(http.StatusBadRequest).
		WithTitle("Could not parse JSON request body").
		WithDetail("The request body is invalid.")
}
