package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darkhunterLearning/Go-CRUD-API-JWT/api"
	"github.com/darkhunterLearning/Go-CRUD-API-JWT/middleware"
	"github.com/gorilla/mux"
)

func Init_Route() {
	router := mux.NewRouter()

	router.HandleFunc("/signup/", api.SignUp).Methods("POST")
	router.HandleFunc("/login/", api.LogIn).Methods("POST")
	// router.HandleFunc("/customers/", api.GetCustomers).Methods("GET")
	router.Handle("/customers/", middleware.CheckJwt(http.HandlerFunc(api.GetCustomers))).Methods("GET")
	router.Handle("/customer/{id}", middleware.CheckJwt(http.HandlerFunc(api.GetCustomer))).Methods("GET")

	fmt.Println("Listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
