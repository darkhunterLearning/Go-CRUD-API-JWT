package api

import (
	"fmt"
	"net/http"

	"github.com/darkhunterLearning/Go-CRUD-API-JWT/action"
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/db"
	"github.com/gorilla/mux"
)

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ID := vars["id"]

	if ID == "" {
		action.JSON(w, 401, "Missing ID")
	} else {
		db := db.ConnectDB()

		fmt.Println("Deleting customer from DB")

		_, err := db.Exec("DELETE FROM customer where id = $1", ID)

		// check errors
		action.CheckErr(err)

		action.JSON(w, 200, "Customer has been deleted successfully!")
	}
}
