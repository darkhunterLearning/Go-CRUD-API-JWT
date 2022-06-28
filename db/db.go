package db

import (
	"fmt"
	"log"

	"github.com/darkhunterLearning/Go-CRUD-API-JWT/model"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func GetDatabase() *gorm.DB {
	username := "postgres"
	password := "12345"
	dbName := "Customer"
	dbHost := "localhost"

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUri)
	connection, err := gorm.Open("postgres", dbUri)
	if err != nil {
		// fmt.Println(err)
		log.Fatalln("wrong database url")
	}

	sqldb := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("database connected")
	}

	fmt.Println("connected to database")
	return connection
}
func InitialMigration() {
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.AutoMigrate(model.Customer{})
}

func CloseDatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}
