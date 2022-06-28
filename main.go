package main

import (
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/api"
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/db"
)

func main() {
	db.InitialMigration()
	api.CreateRouter()
	api.InitializeRoute()
	api.ServerStart()
}
