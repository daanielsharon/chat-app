package app

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	dataSourceName := "postgresql://root:root@localhost:1234/go-realtimechat?sslmode=disable"
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		return nil, err
	}
	
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}