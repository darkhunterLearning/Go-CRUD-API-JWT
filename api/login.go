package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/action"
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/db"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	customer_email := r.PostFormValue("customer_email")
	customer_password := r.PostFormValue("customer_password")

	if govalidator.IsNull(customer_email) || govalidator.IsNull(customer_password) {
		action.JSON(w, 400, "Data can not empty")
		return
	}

	if !govalidator.IsEmail(customer_email) {
		action.JSON(w, 400, "Email is invalid")
		return
	}

	customer_email = action.Santize(customer_email)
	customer_password = action.Santize(customer_password)

	postgres_db := db.ConnectDB()

	findEmail := `SELECT customer_email, customer_password FROM customer WHERE customer_email=$1;`

	row := postgres_db.QueryRow(findEmail, customer_email)

	var email string
	var password string

	switch err := row.Scan(&email, &password); err {
	case sql.ErrNoRows:
		action.JSON(w, 400, "Username or Password incorrect")
	case nil:
		hashedPassword := fmt.Sprintf("%v", password)
		err = action.CheckPasswordHash(hashedPassword, customer_password)

		if err != nil {
			action.JSON(w, 401, "Username or Password incorrect")
			return
		}
		token, errCreate := action.Create(customer_email)

		if errCreate != nil {
			action.JSON(w, 500, "Internal Server Error")
			return
		}

		action.JSON(w, 200, token)
	default:
		panic(err)
	}
}
