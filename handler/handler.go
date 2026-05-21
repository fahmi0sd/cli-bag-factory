package handler

import (
	"database/sql"

	"github.com/fahmi0sd/cli-bag-factory/cli"
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

// function to start the CLI application
func (h *CLIHandler) Start() {
	cli.InterfaceCLI()
}
