package todo

import (
	"database/sql"
	"fmt"
)

// Project a project for assigning tasks to
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
}

// Insert insert a new task to the backend
func (t Task) Insert(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	stmt, err := db.Prepare("select id from projects where name = ?")
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	var i int
	err = stmt.QueryRow(t.Project.Name).Scan(&i)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	if err == sql.ErrNoRows {
		res, err := tx.Exec("insert into projects (name) values (?)", t.Project.Name)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return err
		}
		last, err := res.LastInsertId()
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return err
		}
		i = int(last)

	}
	t.Project.ID = int(i)

	_, err = tx.Exec("insert into tasks (name, priority, completed, project_id) values (?,?,?,?)",
		t.Name, t.Priority, t.Completed, t.Project.ID)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// List fetch and return open tasks from the backend
func List(db *sql.DB) ([]Task, error) {
	rows, err := db.Query(`
	select t.id, t.name, priority, coalesce(p.id,0) as project_id, coalesce(p.name, "") as project_name
	from tasks t
	left outer join projects p on p.id = t.project_id
	where completed = 0 
	order by priority asc, t.created asc`)
	if err != nil {
		fmt.Println(err)
		return []Task{}, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err = rows.Scan(&t.ID, &t.Name, &t.Priority, &t.Project.ID, &t.Project.Name)
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
