package doit

import (
	"database/sql"
	"fmt"
)

// Project a project to hold releated tasks
type Project struct {
	ID       int
	Name     string
	Created  int
	Modified int
}

// Task a single todo task
type Task struct {
	ID        int
	Name      string
	Priority  int
	Completed int
	Project   Project
	Created   int
	Modified  int
	DB        *sql.DB
}

// Insert insert a new task to the backend
func (t Task) Insert() error {
	tx, err := t.DB.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = tx.Exec("insert into tasks (name, priority, completed) values (?,?,?)",
		t.Name, t.Priority, t.Completed)
	if err != nil {
		fmt.Println(err)
		return err
	}

	tx.Commit()
	return nil
}

// List fetch and return open tasks from the backend
func List(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("select id, name, priority from tasks where completed = 0 order by priority asc, created asc")
	if err != nil {
		fmt.Println(err)
		return []Task{}, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err = rows.Scan(&t.ID, &t.Name, &t.Priority)
		if err != nil {
			fmt.Println(err)
			return []Task{}, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

// Done mark a task as completed
func Done(id int, db *sql.DB) error {
	_, err := db.Exec("update tasks set completed = 1 where id = ?", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
