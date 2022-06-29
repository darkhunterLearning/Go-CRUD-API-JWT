package api

import (
	"database/sql"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/action"
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/db"
	"github.com/google/uuid"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	customer_name := r.PostFormValue("customer_name")
	customer_phone := r.PostFormValue("customer_phone")
	customer_address := r.PostFormValue("customer_address")
	customer_email := r.PostFormValue("customer_email")
	customer_password := r.PostFormValue("customer_password")
	// fmt.Println(customer_name, customer_phone, customer_address, customer_email, customer_password)
	if govalidator.IsNull(customer_name) || govalidator.IsNull(customer_phone) || govalidator.IsNull(customer_address) || govalidator.IsNull(customer_email) || govalidator.IsNull(customer_password) {
		action.JSON(w, 400, "Data can not empty")
		return
	}

	if !govalidator.IsEmail(customer_email) {
		action.JSON(w, 400, "Email is invalid")
		return
	}
	customer_uniqueID := uuid.New().String()
	customer_name = action.Santize(customer_name)
	customer_phone = action.Santize(customer_phone)
	customer_address = action.Santize(customer_address)
	customer_email = action.Santize(customer_email)
	customer_password = action.Santize(customer_password)

	postgres_db := db.ConnectDB()

	checkExistedEmail := `SELECT customer_email FROM customer WHERE customer_email=$1;`

	row := postgres_db.QueryRow(checkExistedEmail, customer_email)
	var email string

	switch err := row.Scan(&email); err {
	case sql.ErrNoRows:
		password, err := action.Hash(customer_password)

		if err != nil {
			action.JSON(w, 500, "Register has failed")
			return
		}
		signUpSQL := `INSERT INTO customer(unique_id, customer_name, customer_phone, customer_address, customer_email, customer_password)
					 VALUES($1, $2, $3, $4, $5, $6)`
		_, err = postgres_db.Exec(signUpSQL, customer_uniqueID, customer_name, customer_phone, customer_address, customer_email, password)

		if err != nil {
			action.JSON(w, 500, "Register has failed")
		}

		action.JSON(w, 201, "Register Succesfully")
	case nil:
		action.JSON(w, 409, "User does exists")
	default:
		panic(err)
	}
}
