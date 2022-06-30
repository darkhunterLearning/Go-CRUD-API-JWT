package api

import (
	"fmt"
	"net/http"

	"github.com/darkhunterLearning/Go-CRUD-API-JWT/action"
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/db"
)

func DeleteCustomers(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectDB()

	fmt.Println("Deleting all customer!")

	_, err := db.Exec("DELETE FROM customer")

	action.CheckErr(err)

	fmt.Println("All customers have been deleted successfully")

	action.JSON(w, 200, "All customers have been deleted successfully")
}
