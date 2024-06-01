package marktodoasdone

func Handle(input Input, dataStore Store) error {
	storeTodo, err := dataStore.GetTodoByID(input.TodoID)

	if err != nil {
		return err
	}

	domainTodo, err := storeTodo.ToDomainTodo()

	if err != nil {
		return err
	}

	err = domainTodo.MarkAsDone()

	if err != nil {
		return err
	}

	err = dataStore.MarkTodoAsDone(input.TodoID)

	if err != nil {
		return nil
	}

	return err
}
