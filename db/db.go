package db

import (
	"database/sql"
	"fmt"

	"github.com/darkhunterLearning/Go-CRUD-API-JWT/action"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "12345"
	DB_NAME     = "Customer"
)

func ConnectDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	DB, err := sql.Open("postgres", dbinfo)
	action.CheckErr(err)

	return DB
}
