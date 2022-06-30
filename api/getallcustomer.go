package api

import (
	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/darkhunterLearning/Go-CRUD-API-JWT/action"
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/db"
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/model"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectDB()

	fmt.Println("Getting customers...")

	// Get all customers from customer table
	rows, err := db.Query("SELECT * FROM customer")

	// check errors
	action.CheckErr(err)

	// var response []JsonResponse
	var customers []model.Customer

	// Foreach customer
	for rows.Next() {
		var (
			id               uint32
			uniqueID         string
			customerName     string
			customerPhone    string
			customerAddress  string
			customerEmail    string
			customerPassword string
		)

		err = rows.Scan(&id, &uniqueID, &customerName, &customerPhone, &customerAddress, &customerEmail, &customerPassword)

		// check errors
		action.CheckErr(err)

		customers = append(customers, model.Customer{ID: id, Unique_Id: uniqueID, Customer_Name: customerName, Customer_Phone: customerPhone,

			Customer_Address: customerAddress, Customer_Email: customerEmail, Customer_Password: customerPassword})
	}

	action.JSON(w, 200, customers)

}
