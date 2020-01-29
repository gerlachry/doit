package doit

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

// Insert insert a new project to the backend
func (p *Project) Insert(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	res, err := tx.Exec("insert into projects (name) values (?)", p.Name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	i, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return err
	}
	p.ID = int(i)
	tx.Commit()
	return nil
}

// Find lookup a project
func Find(name string, db *sql.DB) (Project, error) {
	stmt, err := db.Prepare("select id, name, created, modified from projects where name = ?")
	if err != nil {
		fmt.Println(err)
		return Project{}, err
	}
	defer stmt.Close()

	var p Project
	err = stmt.QueryRow(name).Scan(&p.ID, &p.Name, &p.Created, p.Modified)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return Project{}, err
	}

	return p, nil
}
