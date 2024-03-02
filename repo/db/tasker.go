package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

const (
	createProjectTableQuery = `
	CREATE TABLE IF NOT EXISTS Projects (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title text NOT NULL,
	description text,
	priority text CHECK( priority IN ('Low', 'Medium', 'High') ) NOT NULL DEFAULT 'Low'
);`
	createTaskTableQuery = `
	CREATE TABLE IF NOT EXISTS Tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	projectId INTEGER NOT NULL REFERENCES Projects(id),
	title text NOT NULL,
	description text,
	status text CHECK ( status IN ('Pending', 'InProgress', 'Done')) DEFAULT 'Pending'
);`
)

type SqliteDB struct {
	DB      *sql.DB
	dataDir string
}

func Open(path string) (*SqliteDB, error) {
	db, err := sql.Open("sqlite", fmt.Sprintf("%s/tasker.db", path))
	if err != nil {
		return nil, err
	}
	log.Println("Database opened at path:", path)
	tables := []string{"Projects", "Tasks"}
	repo := &SqliteDB{db, path}
	for _, table := range tables {
		exist, err := repo.tableExists(table)
		if err != nil {
			return nil, err
		}
		if !exist {
			err = repo.createTable(table)
			if err != nil {
				return nil, err
			}
		}
	}
	return repo, err
}

func (sq *SqliteDB) tableExists(tableName string) (bool, error) {
	var table string
	err := sq.DB.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name=?", tableName).Scan(&table)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("Table %s does not exist", tableName)
		return false, nil
	case err != nil:
		log.Println("Error checking if table exists: ", err)
		return false, err
	default:
		log.Printf("Table %s exists", tableName)
		return true, nil
	}
}

func (sq *SqliteDB) createTable(tableName string) error {
	var tableQuery string
	switch tableName {
	case "Projects":
		tableQuery = createProjectTableQuery
	case "Tasks":
		tableQuery = createTaskTableQuery
	default:
		return fmt.Errorf("table %s not found", tableName)
	}
	_, err := sq.DB.Exec(tableQuery)
	if err != nil {
		log.Println("Error creating table: ", err)
		return err
	}
	log.Printf("Table %s created", tableName)
	return nil
}
