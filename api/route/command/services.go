package command

import (
	"github.com/matthiashermsen/dewit/domain/command/marktodoasdone"
	"github.com/matthiashermsen/dewit/domain/command/notetodo"
)

type Services struct {
	NoteTodo       notetodo.ServiceFunc
	MarkTodoAsDone marktodoasdone.ServiceFunc
}
