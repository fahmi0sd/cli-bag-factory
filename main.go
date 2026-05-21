package main

import (
	"log"

	"github.com/fahmi0sd/cli-bag-factory/cli"
	"github.com/fahmi0sd/cli-bag-factory/db"
	"github.com/fahmi0sd/cli-bag-factory/handler"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// database connection
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer database.Close()

	// handler untuk CLI
	cliHandler := handler.NewCLIHandler(database)

	// Oper handler CLI
	cli.InterfaceCLI(cliHandler)
}
