package api

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/action"
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/db"
	"github.com/gorilla/mux"
)

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	if ID == "" {
		action.JSON(w, 401, "You are missing ID parameter.")
	} else {
		db := db.ConnectDB()
		Customer_Name := r.FormValue("customer_name")
		Customer_Phone := r.FormValue("customer_phone")
		Customer_Address := r.FormValue("customer_address")
		Customer_Email := r.FormValue("customer_email")
		Customer_Password := r.FormValue("customer_password")

		if govalidator.IsNull(Customer_Name) || govalidator.IsNull(Customer_Phone) || govalidator.IsNull(Customer_Address) || govalidator.IsNull(Customer_Email) || govalidator.IsNull(Customer_Password) {
			action.JSON(w, 400, "Data can not empty")
			return
		}

		if !govalidator.IsEmail(Customer_Email) {
			action.JSON(w, 400, "Email is invalid")
			return
		}

		sqlUpdate := `
		UPDATE customer
		SET customer_name = $1, customer_phone = $2, customer_address = $3, customer_email = $4, customer_password = $5
		WHERE id = $6;`
		_, err := db.Exec(sqlUpdate, Customer_Name, Customer_Phone, Customer_Address, Customer_Email, Customer_Password, ID)

		action.CheckErr(err)

		action.JSON(w, 200, "Customer's Info has been updated successfully!")
	}
}
