package helper

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// function to setup test database connection in unit test
func SetupTestDB() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("⚠️ File .env tidak ditemukan, menggunakan variabel sistem default")
	}

	dsn := os.Getenv("MYSQL_DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Gagal koneksi ke test database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("❌ Gagal ping ke test database: %v", err)
	}

	return db
}
