package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

var DB *sql.DB

func InitDB() {
	var err error

	// FIXME: no .env variables available
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_addr := db_user + ":" + db_password + "@/" + db_name

	fmt.Println("initializing DB:", db_addr)

	DB, err = sql.Open("mysql", db_addr)
	if err != nil {
		panic(err)
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
