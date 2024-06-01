package notetodo

type Input struct {
	TodoTitle string
}

func NewInput(todoTitle string) Input {
	return Input{
		TodoTitle: todoTitle,
	}
}
