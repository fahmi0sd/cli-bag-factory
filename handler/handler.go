package handler

import (
	"database/sql"
)

// Struct for database connection and handler methods
type CLIHandler struct {
	DB *sql.DB
}

// function to create a new CLIHandler with database connection
func NewCLIHandler(db *sql.DB) *CLIHandler {
	return &CLIHandler{
		DB: db,
	}
}
