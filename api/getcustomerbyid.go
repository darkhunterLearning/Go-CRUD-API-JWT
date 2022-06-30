package api

import (
	"fmt"
	"net/http"

	"github.com/darkhunterLearning/Go-CRUD-API-JWT/action"
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/db"
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/model"
	"github.com/gorilla/mux"
)

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	// var response = JsonResponse{}
	db := db.ConnectDB()

	fmt.Println("Getting customer by id...")

	// Get all customers from customer table
	rows, err := db.Query("SELECT * FROM customer WHERE id = $1;", ID)

	// check errors
	action.CheckErr(err)

	// var response []JsonResponse
	var customer []model.Customer

	// Foreach customer
	for rows.Next() {
		var (
			ID               uint32
			uniqueID         string
			customerName     string
			customerPhone    string
			customerAddress  string
			customerEmail    string
			customerPassword string
		)

		err = rows.Scan(&ID, &uniqueID, &customerName, &customerPhone, &customerAddress, &customerEmail, &customerPassword)

		// check errors
		action.CheckErr(err)

		customer = append(customer, model.Customer{ID: ID, Unique_Id: uniqueID, Customer_Name: customerName, Customer_Phone: customerPhone,

			Customer_Address: customerAddress, Customer_Email: customerEmail, Customer_Password: customerPassword})
	}

	action.JSON(w, 200, customer)

}
