package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// ⚠️ Basic validation (very important)
	if dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("❌ Database environment variables not set properly")
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		dbUser,
		dbPassword,
		dbName,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Failed to open DB connection:", err)
	}

	// 🔥 Verify connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}

	fmt.Println("✅ Connected to DB")
}