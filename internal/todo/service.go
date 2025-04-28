package todo

func Create(userID int, title string) error {
	return CreateTodo(userID, title)
}

func List(userID int) ([]Todo, error) {
	return GetTodos(userID)
}

func Update(id int, title string, done bool) error {
	return UpdateTodoStatus(id, title, done)
}

func Delete(id int) error {
	return DeleteTodo(id)
}
