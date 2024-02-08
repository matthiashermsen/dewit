package response

func GetInternalServerErrorProblemDetails() ProblemDetails {
	return NewProblemDetails().
		WithInternalError()
}
