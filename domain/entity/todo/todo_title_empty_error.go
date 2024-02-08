package todo

type TodoTitleEmptyError struct{}

func (e *TodoTitleEmptyError) Error() string {
	return "Todo title must not be empty."
}
