package models

import "github.jcosta86.com/todoapi/db"

// Insert inserts a new todo
func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING id`
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Completed).Scan(&id)
	return
}

// Get returns a todo by id
func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT id, title, description, completed FROM todos WHERE id = $1`
	row := conn.QueryRow(sql, id)
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
	if err != nil {
		return
	}

	return
}

// GetAll returns all todos
func GeAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

// Update updates a todo
func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := `UPDATE todos SET title = $1, description = $2, completed = $3 WHERE id = $4`
	res, err := conn.Exec(sql, todo.Title, todo.Description, todo.Completed, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Delete deletes a todo
func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := `DELETE FROM todos WHERE id = $1`
	res, err := conn.Exec(sql, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
