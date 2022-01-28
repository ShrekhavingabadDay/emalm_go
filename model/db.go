package model

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	err := godotenv.Load("config/.env")

	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err.Error())
		return
	}

	connInfo := mysql.NewConfig()

	connInfo.User = os.Getenv("DB_USER")
	connInfo.Passwd = os.Getenv("DB_PASSWORD")
	connInfo.Net = "tcp"
	connInfo.Addr = "127.0.0.1:3306"
	connInfo.DBName = os.Getenv("DB_NAME")

	db, err = sql.Open("mysql", connInfo.FormatDSN())

	if err != nil {
		fmt.Printf("Error connecting to the DB: %s\n", err.Error())
		return
	} else {
		fmt.Printf("Successfully connected to DB!\n")
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Could not ping database: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB is open!\n")
	}

}
