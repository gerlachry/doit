package pkg

import (
	"database/sql"
	"fmt"
)

// Task a single todo task
type Task struct {
	ID        int
	Name      string
	Priority  int
	Completed int
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
