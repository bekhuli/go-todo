package todo

import "github.com/bekhuli/go-todo/internal/db"

func CreateTodo(userID int, title string) error {
	_, err := db.DB.Exec("INSERT INTO todos (user_id, title, done) VALUES ($1, $2, false)", userID, title)
	return err
}

func GetTodos(userID int) ([]Todo, error) {
	rows, err := db.DB.Query("SELECT id, title, done FROM todos WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.Title, &todo.Done)
		todos = append(todos, todo)
	}

	return todos, nil
}

func UpdateTodoStatus(id int, title string, done bool) error {
	_, err := db.DB.Exec("UPDATE todos SET done = $1, title = $2 WHERE id = $3", done, title, id)
	return err
}

func DeleteTodo(id int) error {
	_, err := db.DB.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}
